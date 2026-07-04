package jina2

import (
	"context"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
)

func (f *Jinja2Feature) Hover(ctx context.Context, doc *document.Document, pos protocol.Position) (*protocol.HoverResult, error) {
	return nil, nil
	// tree := f.trees[doc.URI]
	// if tree == nil {
	// 	return nil, nil
	// }
	// node := ast.FindNodeAtPosition(tree.RootNode(), pos.Line, pos.Character)
	// if node == nil {
	// 	return nil, nil
	// }
	// keyword, ok := f.keyword.Lookup(node.Kind())
	// if !ok {
	// 	return nil, nil
	// }
	// return &protocol.HoverResult{
	// 	Contents: protocol.MarkupContent{
	// 		Kind:  keyword.Description.Format,
	// 		Value: keyword.Description.Value,
	// 	},
	// }, nil
}
