package jina2

import (
	"context"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
)

func (f *Jinja2Feature) Definition(ctx context.Context, doc *document.Document, pos protocol.Position) ([]protocol.Location, error) {
	return []protocol.Location{
		{
			URI: doc.URI,
			Range: protocol.Range{
				Start: protocol.Position{Line: pos.Line - 1, Character: 0},
				End:   protocol.Position{Line: pos.Line - 1, Character: 0},
			},
		},
	}, nil
}
