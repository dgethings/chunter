package cisco_ios

import (
	"context"

	ts_ci "github.com/dgethings/chunter/grammars/tree-sitter-cisco_ios/bindings/go"
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

type CiscoIOSFeature struct {
	parser  *sitter.Parser
	trees   map[string]*sitter.Tree
	lang    *sitter.Language
	keyword keyword.Keywords
}

func New() *CiscoIOSFeature {
	lang := sitter.NewLanguage(ts_ci.Language())
	p := sitter.NewParser()
	p.SetLanguage(lang)
	return &CiscoIOSFeature{
		parser:  p,
		trees:   make(map[string]*sitter.Tree),
		lang:    lang,
		keyword: Keywords,
	}
}

func (f *CiscoIOSFeature) LanguageID() string {
	return "cisco_ios"
}

func (f *CiscoIOSFeature) Close() {
	for _, t := range f.trees {
		t.Close()
	}
	f.parser.Close()
}

func (f *CiscoIOSFeature) DidOpen(ctx context.Context, doc *document.Document) error {
	tree := f.parser.Parse(doc.Content, nil)
	if tree != nil {
		f.trees[doc.URI] = tree
	}
	return nil
}

func (f *CiscoIOSFeature) DidChange(ctx context.Context, doc *document.Document) ([]protocol.Diagnostic, error) {
	oldTree := f.trees[doc.URI]
	newTree := f.parser.Parse(doc.Content, oldTree)
	if oldTree != nil {
		oldTree.Close()
	}
	f.trees[doc.URI] = newTree
	return f.runDiagnostics(doc, newTree), nil
}

func (f *CiscoIOSFeature) DidClose(ctx context.Context, doc *document.Document) error {
	if t, ok := f.trees[doc.URI]; ok {
		t.Close()
		delete(f.trees, doc.URI)
	}
	return nil
}
