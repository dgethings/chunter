package cisco_ios_jinja2

import (
	"context"
	"log/slog"
	"runtime/debug"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
	"github.com/dgethings/chunter/internal/symbols"
	ts "github.com/dgethings/tree-sitter-cisco-ios-jinja2/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// CiscoIOSFeature implements the LSP feature set for Cisco IOS + Jinja2 configs.
// All methods are called from a single goroutine (the LSP server runs with
// jrpc2 Concurrency=1), so the mutable fields below (trees, parser) need no
// synchronization. See cmd/serve.go.
type CiscoIOSFeature struct {
	parser  *sitter.Parser
	trees   map[string]*sitter.Tree
	keyword *keyword.Set
	symbols *symbols.Table
}

func New() *CiscoIOSFeature {
	p := sitter.NewParser()
	p.SetLanguage(sitter.NewLanguage(ts.Language()))
	slog.Info("grammar", "module", GrammarModule, "version", GrammarVersion())
	return &CiscoIOSFeature{
		parser:  p,
		trees:   make(map[string]*sitter.Tree),
		keyword: keyword.NewSet(Keywords),
		symbols: symbols.NewTable(),
	}
}

// GrammarModule is the tree-sitter grammar dependency, pulled from go.mod at
// build time. Keep in sync with the require line in go.mod.
const GrammarModule = "github.com/dgethings/tree-sitter-cisco-ios-jinja2"

// GrammarVersion reads the resolved version of the grammar module from the
// binary's build info (populated at `go build`). Returns "" when build info is
// unavailable (e.g. a stripped binary that lacks it).
func GrammarVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	for _, dep := range bi.Deps {
		if dep.Path == GrammarModule {
			return dep.Version
		}
	}
	return ""
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
