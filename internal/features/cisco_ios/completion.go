package cisco_ios

import (
	"context"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
)

func (f *CiscoIOSFeature) Completion(ctx context.Context, doc *document.Document, pos protocol.Position) ([]protocol.CompletionItem, error) {
	return []protocol.CompletionItem{
		{
			Label:         "NAF",
			Detail:        "Not going this year :(",
			Documentation: "First one in EU that I'm missing. V sad",
		},
	}, nil
}
