package server

import (
	"context"
	"log/slog"

	"github.com/creachadair/jrpc2"
	"github.com/dgethings/chunter/internal/document"
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
	l := slog.With("language", doc.LanguageID, "message", "didOpen")
	l.Debug("stored document", "uri", doc.URI, "version", doc.Version)

	f, err := s.features.Route(doc.LanguageID)
	if err != nil {
		l.Error("failed to find supported language", "error", err.Error())
		return nil
	}
	diagnostics, err := f.DidOpen(ctx, doc)
	if err != nil {
		l.Error("didOpen error", "error", err.Error())
	}

	srv := jrpc2.ServerFromContext(ctx)
	if srv != nil {
		srv.Notify(ctx, "textDocument/publishDiagnostics", protocol.PublishDiagnosticsParams{
			URI:         doc.URI,
			Diagnostics: diagnostics,
		})
	}
	l.Debug("opened", "uri", doc.URI)
	return nil
}

func (s *Server) DidChange(ctx context.Context, params protocol.DidChangeTextDocumentParams) error {
	doc, err := s.documents.Get(params.TextDocument.URI)
	if err != nil {
		slog.Error("failed to get document", "error", err.Error())
		return nil
	}
	l := slog.With("message", "didChange", "uri", doc.URI, "language", doc.LanguageID)

	for _, change := range params.ContentChanges {
		doc.Content = []byte(change.Text)
		doc.Version = params.TextDocument.Version
		s.documents.Put(doc)

		f, err := s.features.Route(doc.LanguageID)
		if err != nil {
			l.Error("failed to find supported language", "language", doc.LanguageID, "error", err)
			continue
		}

		diagnostics, err := f.DidChange(ctx, doc)
		if err != nil {
			l.Error("failed to get diagnostics", "language", doc.LanguageID, "error", err)
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
	l.Debug("successfully changed", "uri", doc.URI)
	return nil
}

func (s *Server) DidClose(ctx context.Context, params protocol.TextDocumentIdentifier) error {
	doc, err := s.documents.Get(params.URI)
	if err != nil {
		slog.Error("failed to get document", "error", err.Error())
		return nil
	}
	l := slog.With("uri", doc.URI, "language", doc.LanguageID, "message", "didClose")

	f, err := s.features.Route(doc.LanguageID)
	if err != nil {
		l.Error("failed to find supported language", "error", err.Error())
		return nil
	}
	if err := f.DidClose(ctx, doc); err != nil {
		l.Error("failed execution", "error", err.Error())
	}
	s.documents.Delete(params.URI)
	return nil
}
