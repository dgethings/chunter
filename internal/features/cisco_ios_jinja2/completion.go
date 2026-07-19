package cisco_ios_jinja2

import (
	"context"
	"log/slog"
	"regexp"
	"strings"

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

	// sectionForNode maps a tree-sitter *_section node kind to the
	// corresponding keyword.Section string used in the keywords data.
	// The innermost section in the ancestor chain wins (break on first
	// match walking up).
	//
	// Gaps: the keyword data carries ~600 distinct Section values
	// (config-pmap-c, config-archive, config-vpdn, …) for which the
	// grammar does not yet emit a *_section node. Those keywords remain
	// unreachable from completion until the grammar is extended.
	sectionForNode := map[string]string{
		"interface_section":      "config-if",
		"router_section":         "config-router",
		"route_map_section":      "config-route-map",
		"class_map_section":      "config-cmap",
		"policy_map_section":     "config-pmap",
		"vlan_section":           "config-vlan",
		"line_section":           "config-line",
		"ip_access_list_section": "config-ext-nacl",
	}
	section := "config"
	for n := node; n != nil; n = n.Parent() {
		if s, ok := sectionForNode[n.Kind()]; ok {
			section = s
			break
		}
	}

	items := createItems(f.keyword.InSection(section))
	return items, nil
}

// placeholderDefaultRe matches LSP snippet placeholders of the form ${N:default}
// and strips the default text, leaving an empty tabstop ${N}.
//
// Neovim's vim.snippet selects non-empty tabstops via a key sequence that some
// completion engines (notably blink.cmp on Neovim 0.12) leave in INSERT mode
// rather than SELECT mode, so typing at the placeholder inserts before the
// default text instead of replacing it (e.g. `hostname ${1:name}` yielded
// `hostname r1name` instead of `hostname r1`). Empty tabstops use a different
// code path that drops the cursor in INSERT mode at the tabstop position, so
// typing inserts at the correct spot.
var placeholderDefaultRe = regexp.MustCompile(`\$\{(\d+):[^}]*\}`)

// stripPlaceholderDefaults rewrites a snippet by removing the default text from
// every ${N:default} placeholder while leaving the tabstop itself intact.
func stripPlaceholderDefaults(snippet string) string {
	return placeholderDefaultRe.ReplaceAllString(snippet, "${$1}")
}

func createItems(kws keyword.Keywords) []protocol.CompletionItem {
	items := []protocol.CompletionItem{}
	format := protocol.InsertTextFormatSnippet
	kind := protocol.CompletionItemKindKeyword
	seen := map[string]bool{}
	for _, kw := range kws {
		for _, s := range kw.Snippets {
			label := snippetLabel(s)
			if seen[label] {
				continue
			}
			seen[label] = true
			snippet := stripPlaceholderDefaults(s)
			filterText := label
			var doc any
			if kw.Description.Value != "" {
				doc = protocol.MarkupContent{
					Kind:  kw.Description.Format,
					Value: kw.Description.Value,
				}
			}
			items = append(items, protocol.CompletionItem{
				Label:            label,
				Documentation:    doc,
				FilterText:       &filterText,
				InsertText:       &snippet,
				InsertTextFormat: &format,
				Kind:             &kind,
			})
		}
	}
	return items
}

// snippetLabel derives a completion item's label from one snippet by taking
// the literal command text that precedes the first ${N:...} placeholder and
// trimming trailing whitespace. Snippets that share a keyword but differ in
// command form — e.g. `activation-character ${1:n}` and `no activation-character`
// — thus yield distinct labels (`activation-character` and
// `no activation-character`) instead of collapsing onto the same one.
func snippetLabel(snippet string) string {
	if i := strings.Index(snippet, "${"); i >= 0 {
		return strings.TrimRight(snippet[:i], " \t")
	}
	return snippet
}
