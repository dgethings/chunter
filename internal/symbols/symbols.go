// Package symbols extracts named Cisco IOS definition sites (interface,
// router, route-map, class-map, policy-map, vlan, line, redundancy, ACL)
// from a parsed tree-sitter AST and indexes them per-URI for downstream
// LSP features (Definition, References, DocumentSymbol, diagnostics).
package symbols

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// Kind classifies a Cisco IOS named definition.
type Kind string

const (
	KindInterface  Kind = "interface"
	KindRouter     Kind = "router"
	KindRouteMap   Kind = "route-map"
	KindClassMap   Kind = "class-map"
	KindPolicyMap  Kind = "policy-map"
	KindVlan       Kind = "vlan"
	KindLine       Kind = "line"
	KindRedundancy Kind = "redundancy"
	KindACL        Kind = "acl"
)

// Symbol is a named definition site extracted from a parsed document.
type Symbol struct {
	Kind      Kind
	Name      string
	URI       string
	Range     protocol.Range // full section / line range
	NameRange protocol.Range // just the name token (Definition / highlight)
}

// Table is a per-URI index of Symbol definitions. Lookups are read-only and
// safe for concurrent use; Index/Clear mutate and must not run concurrently
// with each other on the same URI (the LSP server already serializes per
// document via didOpen/didChange).
type Table struct {
	mu    sync.RWMutex
	byURI map[string][]Symbol
}

func NewTable() *Table {
	return &Table{byURI: make(map[string][]Symbol)}
}

// Index walks root, extracts every definition it recognizes, and replaces
// the stored entry for uri. Call with root=nil to clear.
func (t *Table) Index(uri string, root *sitter.Node, content []byte) {
	syms := Extract(uri, root, content)
	t.mu.Lock()
	t.byURI[uri] = syms
	t.mu.Unlock()
}

// Clear removes all symbols for uri.
func (t *Table) Clear(uri string) {
	t.mu.Lock()
	delete(t.byURI, uri)
	t.mu.Unlock()
}

// All returns every symbol indexed for uri, in document order. The returned
// slice is a copy and may be modified freely.
func (t *Table) All(uri string) []Symbol {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return append([]Symbol(nil), t.byURI[uri]...)
}

// Lookup returns symbols matching kind+name in uri.
func (t *Table) Lookup(uri string, kind Kind, name string) []Symbol {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var out []Symbol
	for _, s := range t.byURI[uri] {
		if s.Kind == kind && s.Name == name {
			out = append(out, s)
		}
	}
	return out
}

// LookupAny returns symbols in uri with a matching name regardless of kind.
func (t *Table) LookupAny(uri string, name string) []Symbol {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var out []Symbol
	for _, s := range t.byURI[uri] {
		if s.Name == name {
			out = append(out, s)
		}
	}
	return out
}

// Extract walks root and returns every symbol it recognizes, in document
// order. Returns nil if root is nil.
//
// Symbols come from two sources:
//   - Hierarchical sections (interface_section, router_section,
//     route_map_section, class_map_section, policy_map_section, vlan_section,
//     line_section, redundancy_section) — each contributes one Symbol via
//     the header's `name` field (or a synthesized name for line/redundancy).
//   - Flat ACL definitions: `ip access-list <standard|extended> NAME` and
//     the numbered form `access-list <N> permit|deny ...` (the latter
//     parses as `command_line(access) + text(-list ...)`; a regex recovers
//     the number from the text node).
func Extract(uri string, root *sitter.Node, content []byte) []Symbol {
	if root == nil {
		return nil
	}
	var out []Symbol
	walkNamed(root, func(n *sitter.Node) bool {
		if sym, ok := extractSection(uri, n, content); ok {
			out = append(out, sym)
			// Section bodies do not themselves contain nested section
			// definitions in IOS, but descending is harmless; keep going
			// so jinja loops containing sections (e.g. `{% for %}` emitting
			// interfaces) are still indexed.
		}
		return true
	})
	walkNamed(root, func(n *sitter.Node) bool {
		if sym, ok := extractACL(uri, n, content); ok {
			out = append(out, sym)
		}
		return true
	})
	return out
}

// sectionSpec describes how to extract a Symbol from a *_section node.
type sectionSpec struct {
	sectionKind string // AST node kind, e.g. "interface_section"
	kind        Kind   // LSP-facing symbol kind
	headerKind  string // child node kind holding the name
	nameField   string // field name on the header carrying the name token
	// synthesize, when non-nil, builds the Symbol.Name from the header node
	// (used by line_section and redundancy_section which have no single
	// `name` field).
	synthesize func(header *sitter.Node, content []byte) (name string, nameRange protocol.Range)
}

var sectionSpecs = []sectionSpec{
	{sectionKind: "interface_section", kind: KindInterface, headerKind: "interface_header", nameField: "name"},
	{sectionKind: "router_section", kind: KindRouter, headerKind: "router_header", nameField: "process_id"},
	{sectionKind: "route_map_section", kind: KindRouteMap, headerKind: "route_map_header", nameField: "name"},
	{sectionKind: "class_map_section", kind: KindClassMap, headerKind: "class_map_header", nameField: "name"},
	{sectionKind: "policy_map_section", kind: KindPolicyMap, headerKind: "policy_map_header", nameField: "name"},
	{sectionKind: "vlan_section", kind: KindVlan, headerKind: "vlan_header", nameField: "name"},
	{
		sectionKind: "line_section",
		kind:        KindLine,
		headerKind:  "line_header",
		synthesize:  synthesizeLineName,
	},
	{
		sectionKind: "redundancy_section",
		kind:        KindRedundancy,
		headerKind:  "redundancy_header",
		synthesize: func(header *sitter.Node, content []byte) (string, protocol.Range) {
			// Singleton — name is fixed; range covers the whole header so
			// Go-To-Definition lands on the `redundancy` keyword.
			return "redundancy", nodeRange(header)
		},
	},
}

func extractSection(uri string, n *sitter.Node, content []byte) (Symbol, bool) {
	for _, sp := range sectionSpecs {
		if n.Kind() != sp.sectionKind {
			continue
		}
		header := namedChildByKind(n, sp.headerKind)
		if header == nil {
			return Symbol{}, false
		}
		sym := Symbol{
			Kind:  sp.kind,
			URI:   uri,
			Range: nodeRange(n),
		}
		if sp.synthesize != nil {
			name, nameRange := sp.synthesize(header, content)
			sym.Name = name
			sym.NameRange = nameRange
			return sym, true
		}
		nameNode := header.ChildByFieldName(sp.nameField)
		if nameNode == nil {
			return Symbol{}, false
		}
		sym.Name = textOf(nameNode, content)
		sym.NameRange = nodeRange(nameNode)
		return sym, true
	}
	return Symbol{}, false
}

// synthesizeLineName builds a stable key from `line <type> <first> [last]`,
// e.g. "vty-0-4" or "console-0". The NameRange covers the whole header
// because there is no single token that names the section.
func synthesizeLineName(header *sitter.Node, content []byte) (string, protocol.Range) {
	typeNode := header.ChildByFieldName("type")
	if typeNode == nil {
		return "", nodeRange(header)
	}
	parts := []string{textOf(typeNode, content)}
	for i := uint(0); i < header.NamedChildCount(); i++ {
		c := header.NamedChild(i)
		if c == nil {
			continue
		}
		// Skip the type field (already captured) and the line_kw keyword.
		if c.Kind() != "value" && c.Kind() != "output" {
			continue
		}
		// The type field itself is also a value/output; skip it by byte range.
		if c.StartByte() == typeNode.StartByte() {
			continue
		}
		parts = append(parts, textOf(c, content))
	}
	return strings.Join(parts, "-"), nodeRange(header)
}

// numberedACLRe matches the `text` companion of a flat numbered ACL line.
// `access-list 101 permit ip any any` parses as
// `command_line(access) + text("-list 101 permit ip any any")`. The regex
// captures the ACL number after the leading `-list`.
var numberedACLRe = regexp.MustCompile(`^-list\s+(\S+)`)

func extractACL(uri string, n *sitter.Node, content []byte) (Symbol, bool) {
	if n.Kind() != "command_line" {
		return Symbol{}, false
	}
	nameNode := n.ChildByFieldName("name")
	if nameNode == nil {
		return Symbol{}, false
	}
	leadingIdent := textOf(nameNode, content)

	// Named ACL: `ip access-list <standard|extended> NAME`.
	// Parses as command_line(ip) + arg(access-list) + arg(standard|extended)
	// + arg(NAME).
	if leadingIdent == "ip" {
		args := namedArgs(n)
		if len(args) >= 3 && textOf(args[0], content) == "access-list" {
			nameArg := args[2]
			sym := Symbol{
				Kind:      KindACL,
				Name:      textOf(nameArg, content),
				URI:       uri,
				Range:     nodeRange(n),
				NameRange: nodeRange(nameArg),
			}
			return sym, true
		}
		return Symbol{}, false
	}

	// Numbered ACL: `access-list <N> permit|deny ...`.
	// Parses as command_line(access) + sibling text("-list <N> ..."). The
	// text node is a sibling of the command_line at the config level
	// (command_line only owns the leading `access` identifier; the residual
	// `-list ...` falls through to a sibling text node because `access-list`
	// is not a prec-2 keyword — see the grammar comment above the
	// *_statement rich rules).
	if leadingIdent == "access" {
		textSibling := n.NextNamedSibling()
		if textSibling == nil || textSibling.Kind() != "text" {
			return Symbol{}, false
		}
		textContent := textOf(textSibling, content)
		m := numberedACLRe.FindStringSubmatch(textContent)
		if m == nil {
			return Symbol{}, false
		}
		// Reconstruct the full ACL identifier "access-list" + number for the
		// name (so references via `ip access-group 101` resolve correctly).
		// Range covers the command_line (access) + the text node; NameRange
		// covers just the access-list token + number in the text node.
		name := fmt.Sprintf("access-list %s", m[1])
		rangeStart := n.StartPosition()
		rangeEnd := textSibling.EndPosition()
		return Symbol{
			Kind: KindACL,
			Name: name,
			URI:  uri,
			Range: protocol.Range{
				Start: protocol.Position{Line: rangeStart.Row, Character: rangeStart.Column},
				End:   protocol.Position{Line: rangeEnd.Row, Character: rangeEnd.Column},
			},
			NameRange: protocol.Range{
				Start: protocol.Position{Line: rangeStart.Row, Character: rangeStart.Column},
				End:   protocol.Position{Line: rangeEnd.Row, Character: rangeEnd.Column},
			},
		}, true
	}
	return Symbol{}, false
}

// walkNamed depth-first traverses the named-children subtree of n, invoking
// visit on each. If visit returns false, the subtree under that node is
// skipped.
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

func namedChildByKind(n *sitter.Node, kind string) *sitter.Node {
	if n == nil {
		return nil
	}
	for i := uint(0); i < n.NamedChildCount(); i++ {
		if c := n.NamedChild(i); c != nil && c.Kind() == kind {
			return c
		}
	}
	return nil
}

// namedArgs returns the named children of n that carry the "arg" field
// (the trailing args of a command_line, after the leading identifier).
// The leading identifier itself (the "name" field) is excluded.
func namedArgs(n *sitter.Node) []*sitter.Node {
	var out []*sitter.Node
	nameNode := n.ChildByFieldName("name")
	for i := uint(0); i < n.NamedChildCount(); i++ {
		c := n.NamedChild(i)
		if c == nil {
			continue
		}
		// Skip the leading identifier (the "name" field).
		if nameNode != nil && c.Id() == nameNode.Id() {
			continue
		}
		out = append(out, c)
	}
	return out
}

func textOf(n *sitter.Node, content []byte) string {
	if n == nil {
		return ""
	}
	return string(content[n.StartByte():n.EndByte()])
}

func nodeRange(n *sitter.Node) protocol.Range {
	if n == nil {
		return protocol.Range{}
	}
	start := n.StartPosition()
	end := n.EndPosition()
	return protocol.Range{
		Start: protocol.Position{Line: start.Row, Character: start.Column},
		End:   protocol.Position{Line: end.Row, Character: end.Column},
	}
}
