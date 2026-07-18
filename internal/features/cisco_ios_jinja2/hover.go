package cisco_ios_jinja2

import (
	"context"

	"github.com/dgethings/chunter/internal/ast"
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
)

func (f *CiscoIOSFeature) Hover(ctx context.Context, doc *document.Document, pos protocol.Position) (*protocol.HoverResult, error) {
	tree := f.trees[doc.URI]
	if tree == nil {
		return nil, nil
	}
	node := ast.FindNodeAtPosition(tree.RootNode(), pos.Line, pos.Character)
	if node == nil {
		return nil, nil
	}
	name := node.Kind()
	if name == "identifier" {
		name = string(doc.Content[node.StartByte():node.EndByte()])
	}
	keyword, ok := f.keyword.Lookup(name)
	if !ok {
		return nil, nil
	}
	return &protocol.HoverResult{
		Contents: protocol.MarkupContent{
			Kind:  keyword.Description.Format,
			Value: keyword.Description.Value,
		},
	}, nil
}
