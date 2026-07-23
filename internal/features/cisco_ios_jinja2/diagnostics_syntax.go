package cisco_ios_jinja2

import (
	"fmt"
	"strings"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// runSyntaxDiagnostics surfaces tree-sitter parse-level problems as LSP
// diagnostics. There are two kinds of recovery nodes the parser can emit:
//
//   - ERROR nodes — tokens that could not be incorporated into any rule. These
//     are reported as SeverityError anchored on the offending span.
//   - MISSING nodes — zero-width markers the parser inserts so a rule can
//     complete despite a missing required token (e.g. an unterminated `{{`
//     yields a MISSING `}}`, and a section without its terminating `!` yields a
//     MISSING `eos`). MISSING nodes are reported as SeverityError with a
//     message naming what is missing, EXCEPT for a MISSING `eos` inside a
//     section, which is a common, low-severity condition downgraded to
//     SeverityWarning and anchored on the section header.
//
// The whole pass is gated on root.HasError(), so clean files (the common
// case) pay only a single boolean check.
func (f *CiscoIOSFeature) runSyntaxDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	if tree == nil {
		return nil
	}
	root := tree.RootNode()
	if root == nil || !root.HasError() {
		return nil
	}

	var diags []protocol.Diagnostic
	walkAll(root, func(n *sitter.Node) bool {
		switch {
		case n.IsError():
			diags = append(diags, protocol.Diagnostic{
				Range:    nodeRangeOnStartRow(n, doc.Content),
				Severity: protocol.SeverityError,
				Source:   "chunter",
				Code:     "syntax-error",
				Message:  fmt.Sprintf("syntax error near %q", firstLine(doc.Content, n)),
			})
			// The ERROR node's children are the recovered tokens it
			// swallowed; descending into them would either re-report them or
			// flag valid nodes that happened to be wrapped, so skip its subtree.
			return false
		case n.IsMissing():
			diags = append(diags, missingDiagnostic(n, doc.Content))
			// MISSING nodes are leaves (zero-width); no subtree to skip, but
			// return false for uniformity.
			return false
		default:
			return true
		}
	})
	return diags
}

// missingDiagnostic builds the diagnostic for a MISSING node. A MISSING eos
// whose parent is a section (the only place eos is required) is downgraded to
// a Warning anchored on the section header; every other missing token is an
// Error anchored at the MISSING node's position (extended to end-of-line so
// the editor has something to render).
func missingDiagnostic(n *sitter.Node, content []byte) protocol.Diagnostic {
	if n.Kind() == "eos" {
		if parent := n.Parent(); parent != nil {
			if _, ok := sectionForNodeMap[parent.Kind()]; ok {
				if hdr := sectionHeaderNode(parent); hdr != nil {
					return protocol.Diagnostic{
						Range:    protocol.LineRange(hdr.StartPosition().Row, hdr.StartPosition().Column, hdr.EndPosition().Column),
						Severity: protocol.SeverityWarning,
						Source:   "chunter",
						Code:     "missing-eos",
						Message:  fmt.Sprintf("section %q is missing its terminating %q", nodeText(hdr, content), "!"),
					}
				}
			}
		}
	}
	kind := n.Kind()
	if kind == "" {
		kind = "token"
	}
	return protocol.Diagnostic{
		Range:    nodeRangeOnStartRow(n, content),
		Severity: protocol.SeverityError,
		Source:   "chunter",
		Code:     "missing-" + kind,
		Message:  fmt.Sprintf("missing %q", kind),
	}
}

// nodeRangeOnStartRow returns a single-line range anchored on the node's start
// row. For a normal single-line node this is the node's own span; for a
// multi-line node (an ERROR wrapping several lines) or a zero-width node (a
// MISSING token) the range is extended to the end of the start row so the
// editor always has a non-empty region to underline.
func nodeRangeOnStartRow(n *sitter.Node, content []byte) protocol.Range {
	row := n.StartPosition().Row
	startCol := n.StartPosition().Column
	endCol := n.EndPosition().Column
	if row != n.EndPosition().Row || n.StartByte() == n.EndByte() {
		lineStart := n.StartByte() - startCol
		end := lineStart
		for end < uint(len(content)) && content[end] != '\n' {
			end++
		}
		endCol = end - lineStart
	}
	return protocol.LineRange(row, startCol, endCol)
}

// sectionHeaderNode returns the first named child of a section node, which by
// the grammar is always its *_header node (interface_header, router_header,
// ...). Returns nil if the section has no named children.
func sectionHeaderNode(section *sitter.Node) *sitter.Node {
	if section == nil || section.NamedChildCount() == 0 {
		return nil
	}
	return section.NamedChild(0)
}

// nodeText returns the source slice spanned by n.
func nodeText(n *sitter.Node, content []byte) string {
	if n == nil {
		return ""
	}
	return string(content[n.StartByte():n.EndByte()])
}

// firstLine returns the first line of the node's source text, trimmed of
// surrounding whitespace and truncated so an ERROR spanning many lines does
// not produce a huge diagnostic message.
func firstLine(content []byte, n *sitter.Node) string {
	if n == nil {
		return ""
	}
	s := string(content[n.StartByte():n.EndByte()])
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		s = s[:i]
	}
	s = strings.TrimSpace(s)
	const maxSnippet = 40
	if len(s) <= maxSnippet {
		return s
	}
	return s[:maxSnippet-1] + "…"
}

// walkAll depth-first traverses EVERY child of n (named and anonymous),
// invoking visit on each. This differs from walkNamed: ERROR nodes and
// MISSING anonymous tokens (e.g. a MISSING `}}`) are not named children, so a
// named-only walk would skip them. If visit returns false, the subtree under
// that node is skipped.
func walkAll(n *sitter.Node, visit func(*sitter.Node) bool) {
	if n == nil {
		return
	}
	if !visit(n) {
		return
	}
	for i := uint(0); i < n.ChildCount(); i++ {
		walkAll(n.Child(i), visit)
	}
}
