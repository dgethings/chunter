package server

import (
	"context"

	"github.com/dgethings/chunter/internal/logger"
	"github.com/dgethings/chunter/internal/protocol"
)

func (s *Server) Completion(ctx context.Context, params protocol.CompletionParams) (protocol.CompletionList, error) {
	doc, err := s.documents.Get(params.TextDocument.URI)
	if err != nil {
		logger.FromContext(ctx).Printf("could not find document: %v", err)
		return protocol.CompletionList{}, err
	}
	f, err := s.features.Route(doc.LanguageID)
	if err != nil {
		logger.FromContext(ctx).Printf("could not find language: %s", doc.LanguageID)
		return protocol.CompletionList{}, err
	}
	logger.FromContext(ctx).Println("running completion")
	items, err := f.Completion(ctx, doc, params.Position)
	if err != nil {
		logger.FromContext(ctx).Printf("could not get any completion items: %v", err)
		return protocol.CompletionList{}, err
	}
	logger.FromContext(ctx).Println("success!")
	return protocol.CompletionList{Items: items}, nil
}
