package server

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/logger"
	"github.com/dgethings/chunter/internal/protocol"
)

func (s *Server) DidOpen(ctx context.Context, params protocol.DidOpenTextDocumentParams) error {
	doc := document.New(
		params.TextDocument.URI,
		params.TextDocument.LanguageID,
		params.TextDocument.Version,
		[]byte(params.TextDocument.Text),
	)
	s.documents.Put(doc)

	f, err := s.features.Route(doc.LanguageID)
	if err != nil {
		logger.FromContext(ctx).Printf("didOpen: %s", err)
		return nil
	}
	if err := f.DidOpen(ctx, doc); err != nil {
		logger.FromContext(ctx).Printf("didOpen feature error: %s", err)
	}
	logger.FromContext(ctx).Printf("opened: %s", doc.URI)
	return nil
}

func (s *Server) DidChange(ctx context.Context, params protocol.DidChangeTextDocumentParams) error {
	doc, err := s.documents.Get(params.TextDocument.URI)
	if err != nil {
		logger.FromContext(ctx).Printf("didChange: %s", err)
		return nil
	}

	for _, change := range params.ContentChanges {
		doc.Content = []byte(change.Text)
		doc.Version = params.TextDocument.Version
		s.documents.Put(doc)

		f, err := s.features.Route(doc.LanguageID)
		if err != nil {
			logger.FromContext(ctx).Printf("didChange route: %s", err)
			continue
		}

		diagnostics, err := f.DidChange(ctx, doc)
		if err != nil {
			logger.FromContext(ctx).Printf("didChange feature error: %s", err)
			continue
		}

		srv := jrpc2.ServerFromContext(ctx)
		if srv != nil {
			srv.Notify(ctx, "textDocument/publishDiagnostics", protocol.PublishDiagnosticsParams{
				URI:         doc.URI,
				Diagnostics: diagnostics,
			})
		}
	}
	logger.FromContext(ctx).Printf("changed: %s", doc.URI)
	return nil
}

func (s *Server) DidClose(ctx context.Context, params protocol.TextDocumentIdentifier) error {
	doc, err := s.documents.Get(params.URI)
	if err != nil {
		return nil
	}

	f, err := s.features.Route(doc.LanguageID)
	if err != nil {
		logger.FromContext(ctx).Printf("didClose: %s", err)
		return nil
	}
	if err := f.DidClose(ctx, doc); err != nil {
		logger.FromContext(ctx).Printf("didClose feature error: %s", err)
	}
	s.documents.Delete(params.URI)
	return nil
}
