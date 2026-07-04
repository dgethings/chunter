package jina2

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
	ts_j2 "github.com/dgethings/tree-sitter-jinja2/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

type Jinja2Feature struct {
	parser  *sitter.Parser
	trees   map[string]*sitter.Tree
	lang    *sitter.Language
	keyword keyword.Keywords
}

func New() *Jinja2Feature {
	lang := sitter.NewLanguage(ts_j2.Language())
	p := sitter.NewParser()
	p.SetLanguage(lang)
	if info, err := os.Stat("../tree-sitter-jinja2/src/parser.c"); err == nil {
		slog.Debug("parser version", "jinja2", fmt.Sprintf("%v", info.ModTime()))
	}
	return &Jinja2Feature{
		parser:  p,
		trees:   make(map[string]*sitter.Tree),
		lang:    lang,
		keyword: keyword.Keywords{},
	}
}

func (f *Jinja2Feature) Close() error {
	for _, t := range f.trees {
		t.Close()
	}
	f.trees = make(map[string]*sitter.Tree)
	if f.parser != nil {
		f.parser.Close()
		f.parser = nil
	}
	return nil
}

func (f *Jinja2Feature) LanguageID() string {
	return "jinja2"
}

func (f *Jinja2Feature) Completion(ctx context.Context, doc *document.Document, pos protocol.Position) ([]protocol.CompletionItem, error) {
	return nil, nil
}

func (f *Jinja2Feature) DidOpen(ctx context.Context, doc *document.Document) ([]protocol.Diagnostic, error) {
	tree := f.parser.Parse(doc.Content, nil)
	if tree != nil {
		f.trees[doc.URI] = tree
	}
	// return f.runDiagnostics(doc, tree), nil
	return nil, nil
}

func (f *Jinja2Feature) DidChange(ctx context.Context, doc *document.Document) ([]protocol.Diagnostic, error) {
	oldTree := f.trees[doc.URI]
	newTree := f.parser.Parse(doc.Content, oldTree)
	if oldTree != nil {
		oldTree.Close()
	}
	f.trees[doc.URI] = newTree
	// return f.runDiagnostics(doc, newTree), nil
	return nil, nil
}

func (f *Jinja2Feature) DidClose(ctx context.Context, doc *document.Document) error {
	if t, ok := f.trees[doc.URI]; ok {
		t.Close()
		delete(f.trees, doc.URI)
	}
	return nil
}
