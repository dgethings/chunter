package cisco_ios

import (
	"context"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func (f *CiscoIOSFeature) Hover(ctx context.Context, doc *document.Document, pos protocol.Position) (*protocol.HoverResult, error) {
	tree := f.trees[doc.URI]
	if tree == nil {
		// TODO: log what the doc.URI is, this should never fail. So maybe panic?
		return nil, nil
	}
	keyword, err := nodeByPosition(pos, tree.RootNode())
	if err != nil {
		// TODO: log what the error is and what the node is
		return nil, nil
	}
	return &protocol.HoverResult{
		Contents: protocol.MarkupContent{
			Kind:  keyword.Description.Format,
			Value: keyword.Description.Value,
		},
	}, nil
}

type Keyword struct {
	Description
}

type Description struct {
	Format string
	Value  string
}

func nodeByPosition(pos protocol.Position, root *sitter.Node) (*Keyword, error) {
	pt := sitter.NewPoint(pos.Line, pos.Character)
	node := root.NamedDescendantForPointRange(pt, pt)
	if node == nil {
		return nil, nil
	}
	kw := keywords[node.Kind()]
	return &kw, nil
}

var keywords = map[string]Keyword{
	"hostname_section": {
		Description: Description{
			Format: "plaintext",
			Value:  "To specify or modify the hostname for the network server, use the hostname command in global configuration mode.",
		},
	},
}
