package cisco_ios

import (
	"context"

	"github.com/dgethings/chunter/internal/ast"
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/logger"
	"github.com/dgethings/chunter/internal/protocol"
)

func (f *CiscoIOSFeature) Completion(ctx context.Context, doc *document.Document, pos protocol.Position) ([]protocol.CompletionItem, error) {
	tree := f.trees[doc.URI]
	if tree == nil {
		logger.FromContext(ctx).Printf("cannot find tree for %s\n", doc.URI)
		return nil, nil
	}

	section := ast.FindNodeAtPosition(tree.RootNode(), pos.Line, pos.Character)
	if section == nil {
		logger.FromContext(ctx).Printf("cannot find node for %d:%d\n", pos.Line, pos.Character)
		return nil, nil
	}
	kws := keywords.InSection(section.GrammarName())
	items := []protocol.CompletionItem{}
	for k, i := range kws {
		items = append(items, protocol.CompletionItem{Label: k, Documentation: i.Description.Value})
	}
	return items, nil

	// return []protocol.CompletionItem{
	// {
	// 	Label:         section.Kind(),
	// 	Detail:        fmt.Sprintf("%s:%d:%d", doc.URI, pos.Line, pos.Character),
	// 	Documentation: section.GrammarName(),
	// },
	// }, nil
}
