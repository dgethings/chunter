package ast

import sitter "github.com/tree-sitter/go-tree-sitter"

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
