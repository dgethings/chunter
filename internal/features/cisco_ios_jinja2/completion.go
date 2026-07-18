package cisco_ios_jinja2

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
	slog.Debug("current position", "LINE", pos.Line, "COLOMN", pos.Character)

	node := ast.FindNodeAtPosition(tree.RootNode(), pos.Line, pos.Character)
	if node == nil {
		l.Error("cannot find section", "line", pos.Line, "column", pos.Character)
		return nil, nil
	}
	l.Debug("where am i?", "node", node.GrammarName(), "kind", node.Kind(), "line", pos.Line, "char", pos.Character)

	switch node.Kind() {
	case "value", "text", "eos", "comment", "ios_comment":
		return nil, nil
	}

	section := "config"
	for n := node; n != nil; n = n.Parent() {
		switch n.Kind() {
		case "interface_section":
			section = "config-if"
		case "router_section":
			section = "config-router"
		}
	}

	items := createItems(f.keyword.InSection(section))
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
