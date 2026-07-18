package cisco_ios_jinja2

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
	"github.com/dgethings/chunter/internal/symbols"
	ts "github.com/dgethings/tree-sitter-cisco-ios-jinja2/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

type CiscoIOSFeature struct {
	parser  *sitter.Parser
	trees   map[string]*sitter.Tree
	lang    *sitter.Language
	keyword keyword.Keywords
	symbols *symbols.Table
}

func New() *CiscoIOSFeature {
	lang := sitter.NewLanguage(ts.Language())
	p := sitter.NewParser()
	p.SetLanguage(lang)
	if info, err := os.Stat("../tree-sitter-cisco-ios-jinja2/src/parser.c"); err == nil {
		slog.Info("parser version", "cisco_ios", fmt.Sprintf("%v", info.ModTime()))
	}
	return &CiscoIOSFeature{
		parser:  p,
		trees:   make(map[string]*sitter.Tree),
		lang:    lang,
		keyword: Keywords,
		symbols: symbols.NewTable(),
	}
}

func (f *CiscoIOSFeature) LanguageID() string {
	return "cisco_ios_jinja2"
}

func (f *CiscoIOSFeature) Close() error {
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

func (f *CiscoIOSFeature) DidOpen(ctx context.Context, doc *document.Document) ([]protocol.Diagnostic, error) {
	tree := f.parser.Parse(doc.Content, nil)
	if tree != nil {
		f.trees[doc.URI] = tree
		f.symbols.Index(doc.URI, tree.RootNode(), doc.Content)
	}
	return f.runDiagnostics(doc, tree), nil
}

func (f *CiscoIOSFeature) DidChange(ctx context.Context, doc *document.Document) ([]protocol.Diagnostic, error) {
	oldTree := f.trees[doc.URI]
	newTree := f.parser.Parse(doc.Content, oldTree)
	if oldTree != nil {
		oldTree.Close()
	}
	f.trees[doc.URI] = newTree
	f.symbols.Index(doc.URI, newTree.RootNode(), doc.Content)
	return f.runDiagnostics(doc, newTree), nil
}

func (f *CiscoIOSFeature) DidClose(ctx context.Context, doc *document.Document) error {
	if t, ok := f.trees[doc.URI]; ok {
		t.Close()
		delete(f.trees, doc.URI)
	}
	f.symbols.Clear(doc.URI)
	return nil
}
