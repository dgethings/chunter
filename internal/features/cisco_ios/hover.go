package cisco_ios

import (
	"context"
	"fmt"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
)

func (f *CiscoIOSFeature) Hover(ctx context.Context, doc *document.Document, pos protocol.Position) (*protocol.HoverResult, error) {
	return &protocol.HoverResult{
		Contents: protocol.MarkupContent{
			Kind:  "plaintext",
			Value: fmt.Sprintf("File %s, Characters: %d", doc.URI, len(doc.Content)),
		},
	}, nil
}
