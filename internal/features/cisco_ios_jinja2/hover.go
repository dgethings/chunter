package cisco_ios_jinja2

import (
	"context"
	"strings"

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
	// When the cursor sits exactly at the end of a keyword token,
	// FindNodeAtPosition resolves to the enclosing *_statement / *_header
	// node rather than the anonymous keyword leaf. Recover the keyword text
	// from its first anonymous leaf child so the lookup still matches.
	if strings.HasSuffix(name, "_statement") || strings.HasSuffix(name, "_header") {
		for i := uint(0); i < node.ChildCount(); i++ {
			child := node.Child(i)
			if child == nil {
				continue
			}
			if !child.IsNamed() && child.ChildCount() == 0 {
				name = string(doc.Content[child.StartByte():child.EndByte()])
				break
			}
		}
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
