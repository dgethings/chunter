package cisco_ios

import (
	"context"
	"log/slog"

	"github.com/dgethings/chunter/internal/ast"
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
)

func (f *CiscoIOSFeature) Completion(ctx context.Context, doc *document.Document, pos protocol.Position) ([]protocol.CompletionItem, error) {
	tree := f.trees[doc.URI]
	l := slog.With("uri", doc.URI, "language", doc.LanguageID, "message", "completion")
	if tree == nil {
		l.Error("cannot find tree")
		return nil, nil
	}

	section := ast.FindNodeAtPosition(tree.RootNode(), pos.Line, pos.Character)
	if section == nil {
		l.Error("cannot find section", "line", pos.Line, "column", pos.Character)
		return nil, nil
	}
	l.Info("where am i?", "node", section.GrammarName(), "line", pos.Line, "char", pos.Character)
	kws := f.keyword.InSection(section.GrammarName())
	items := createItems(kws)
	return items, nil
}

func createItems(kws keyword.Keywords) []protocol.CompletionItem {
	items := []protocol.CompletionItem{}
	format := protocol.InsertTextFormatSnippet
	kind := protocol.CompletionItemKindKeyword
	for _, kw := range kws {
		for _, snippet := range kw.Snippets {
			items = append(items, protocol.CompletionItem{
				Label:            kw.Keyword,
				Documentation:    kw.Description.Value,
				InsertText:       &snippet,
				InsertTextFormat: &format,
				Kind:             &kind,
			})
		}
	}
	return items
}
