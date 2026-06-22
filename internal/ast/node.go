package ast

import sitter "github.com/tree-sitter/go-tree-sitter"

// FindNodeAtPosition returns the deepest node (named or anonymous) containing
// the given zero-indexed line and column, or nil if root is nil. Unlike
// NamedNodeAtPosition, this may return anonymous leaf nodes such as literal
// tokens and punctuation.
func FindNodeAtPosition(root *sitter.Node, line, col uint) *sitter.Node {
	if root == nil {
		return nil
	}
	node := root.DescendantForPointRange(
		sitter.Point{Row: line, Column: col},
		sitter.Point{Row: line, Column: col},
	)
	return node
}

func ChildByFieldName(node *sitter.Node, name string) *sitter.Node {
	if node == nil {
		return nil
	}
	return node.ChildByFieldName(name)
}

// NamedChildByKind returns the first direct named child of node whose kind
// matches, or nil if none is found. Useful for reaching named sub-sections
// (e.g. a "service_section") whose fields do not propagate to the parent.
func NamedChildByKind(node *sitter.Node, kind string) *sitter.Node {
	if node == nil {
		return nil
	}
	for i := uint(0); i < node.NamedChildCount(); i++ {
		if c := node.NamedChild(i); c != nil && c.Kind() == kind {
			return c
		}
	}
	return nil
}
