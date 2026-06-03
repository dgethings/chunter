package ast

import sitter "github.com/tree-sitter/go-tree-sitter"

type Tree struct {
	parser *sitter.Parser
	tree   *sitter.Tree
}

func NewTree(lang *sitter.Language) *Tree {
	p := sitter.NewParser()
	p.SetLanguage(lang)
	return &Tree{parser: p}
}

func (t *Tree) Parse(content []byte) *sitter.Node {
	oldTree := t.tree
	t.tree = t.parser.Parse(content, oldTree)
	if oldTree != nil {
		oldTree.Close()
	}
	if t.tree == nil {
		return nil
	}
	return t.tree.RootNode()
}

func (t *Tree) Close() {
	if t.tree != nil {
		t.tree.Close()
	}
	t.parser.Close()
}
