package cisco_ios_jinja2

import (
	"strings"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// runWrongSectionDiagnostics emits a Hint for each command whose keyword is
// known to the database but not valid in the enclosing config section. This
// catches common copy-paste errors (e.g. an interface command inside a
// router section). Unknown keywords are silently skipped — only commands that
// ARE documented but in the WRONG section are flagged.
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

		// Skip negated_statement itself but keep descending so the inner
		// command is still validated where it belongs.
		if kind == "negated_statement" {
			return true
		}

		// Only check command_line and *_statement nodes.
		if kind != "command_line" && !strings.HasSuffix(kind, "_statement") {
			return true
		}

		// Resolve the enclosing section by walking up the ancestors. The ACL
		// special-case is checked before the generic map because
		// ip_access_list_section is present in sectionForNodeMap (mapped to a
		// default) but its true section depends on the header's type field.
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

		// Extract the command keyword from the node's first child. This is the
		// anonymous keyword token for *_statement rules and the named
		// identifier for command_line (matching the leadingAndArgs pattern in
		// the symbols package).
		keyword := firstKeywordFromNode(n, doc.Content)
		if keyword == "" {
			return true
		}

		// Skip unknown keywords (not in the database).
		if _, ok := f.keyword.Lookup(keyword); !ok {
			return true
		}

		// Valid in this section (either an exact match or a global keyword)?
		if f.keyword.IsValidInSection(keyword, enclosingSection) {
			return true
		}

		validSection := f.keyword.LookupSection(keyword)
		diags = append(diags, protocol.Diagnostic{
			Range: protocol.LineRange(
				n.StartPosition().Row,
				n.StartPosition().Column,
				n.EndPosition().Column,
			),
			Severity: protocol.SeverityHint,
			Source:   "chunter",
			Message:  keyword + " is valid in " + validSection + ", not in " + enclosingSection,
		})
		return true
	})

	return diags
}

// firstKeywordFromNode returns the text of the node's first child (the leading
// keyword token). For *_statement rules this is the anonymous keyword literal;
// for command_line it is the named identifier.
func firstKeywordFromNode(n *sitter.Node, content []byte) string {
	if n == nil || n.ChildCount() == 0 {
		return ""
	}
	first := n.Child(0)
	if first == nil {
		return ""
	}
	return string(content[first.StartByte():first.EndByte()])
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
