package cisco_ios_jinja2

import (
	"strings"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// runWrongSectionDiagnostics emits a Hint for each command whose leading
// keyword is known to the database but not valid in the enclosing config
// section. This catches common copy-paste errors (e.g. an interface command
// inside a router section). Unknown keywords are silently skipped - only
// commands that ARE documented but in the WRONG section are flagged.
// For command_line nodes the leading keyword is resolved via longest-prefix
// match against the keyword DB, so multi-word commands like "ip access-group"
// are validated correctly.
func (f *CiscoIOSFeature) runWrongSectionDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	if tree == nil {
		return nil
	}
	root := tree.RootNode()
	if root == nil {
		return nil
	}

	var diags []protocol.Diagnostic

	walkNamed(root, func(n *sitter.Node) bool {
		kind := n.Kind()

		if kind == "negated_statement" {
			return true
		}

		if kind != "command_line" && !strings.HasSuffix(kind, "_statement") {
			return true
		}

		enclosingSection := "config"
		for p := n.Parent(); p != nil; p = p.Parent() {
			if p.Kind() == "ip_access_list_section" {
				enclosingSection = resolveACLSection(p, doc.Content)
				break
			}
			if s, ok := sectionForNodeMap[p.Kind()]; ok {
				enclosingSection = s
				break
			}
		}

		kw := firstKeywordFromNode(n, doc.Content, f.keyword)
		if kw == "" {
			return true
		}

		if _, ok := f.keyword.Lookup(kw); !ok {
			return true
		}

		if f.keyword.IsValidInSection(kw, enclosingSection) {
			return true
		}

		validSection := f.keyword.LookupSection(kw)
		diags = append(diags, protocol.Diagnostic{
			Range: protocol.LineRange(
				n.StartPosition().Row,
				n.StartPosition().Column,
				n.EndPosition().Column,
			),
			Severity: protocol.SeverityHint,
			Source:   "chunter",
			Message:  kw + " is valid in " + validSection + ", not in " + enclosingSection,
		})
		return true
	})

	return diags
}

// firstKeywordFromNode resolves the leading keyword for a command-like
// node. For *_statement rules this is the anonymous first child (single-token
// keyword). For command_line nodes it performs a longest-prefix probe against
// the keyword DB: tokens [name, arg1, arg2, ...] are joined with spaces and
// probed longest-first (up to maxPrefixTokens). This lets multi-word keywords
// like "ip access-group" resolve correctly. Entries containing "(" are
// excluded from probing (documentation aliases, not matchable command text).
// Falls back to the bare name token if no multi-word prefix matches.
func firstKeywordFromNode(n *sitter.Node, content []byte, kw *keyword.Set) string {
	if n == nil || n.ChildCount() == 0 {
		return ""
	}
	first := n.Child(0)
	if first == nil {
		return ""
	}
	name := string(content[first.StartByte():first.EndByte()])

	if n.Kind() != "command_line" || kw == nil {
		return name
	}

	const maxPrefixTokens = 4
	var tokens []string
	tokens = append(tokens, name)
	for i := uint(1); i < n.ChildCount() && len(tokens) < maxPrefixTokens; i++ {
		c := n.Child(i)
		if c == nil || !c.IsNamed() {
			continue
		}
		tokens = append(tokens, string(content[c.StartByte():c.EndByte()]))
	}

	for i := len(tokens); i >= 1; i-- {
		candidate := strings.Join(tokens[:i], " ")
		if entry, ok := kw.Lookup(candidate); ok && !strings.Contains(entry.Keyword, "(") {
			return candidate
		}
	}

	return name
}

// walkNamed depth-first traverses the named-children subtree of n, invoking
// visit on each node. If visit returns false, the subtree under that node is
// skipped. Mirrors the unexported helper in package symbols.
func walkNamed(n *sitter.Node, visit func(*sitter.Node) bool) {
	if n == nil {
		return
	}
	if !visit(n) {
		return
	}
	for i := uint(0); i < n.NamedChildCount(); i++ {
		walkNamed(n.NamedChild(i), visit)
	}
}
