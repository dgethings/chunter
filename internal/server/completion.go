package server

import (
	"context"

	"github.com/dgethings/chunter/internal/protocol"
)

func (s *Server) Completion(ctx context.Context, params protocol.CompletionParams) (protocol.CompletionList, error) {
	doc, err := s.documents.Get(params.TextDocument.URI)
	if err != nil {
		return protocol.CompletionList{}, err
	}
	f, err := s.features.Route(doc.LanguageID)
	if err != nil {
		return protocol.CompletionList{}, err
	}
	items, err := f.Completion(ctx, doc, params.Position)
	if err != nil {
		return protocol.CompletionList{}, err
	}
	return protocol.CompletionList{Items: items}, nil
}
