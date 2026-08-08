package cisco_ios_jinja2

import (
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// The diagnostic passes split into two tiers for progressive publishing
// (chunter-cfz; see DESIGN-chunter-cfz-progressive-diagnostics.md):
//
//   - runTreeDiagnostics: passes that depend ONLY on the parse tree + keyword
//     DB (NOT the symbol table). These can be published as tier 1 immediately
//     after parse, before the expensive symbols.Index.
//   - runRefDiagnostics: the two passes that depend on the symbol table
//     (undefined-refs, duplicate-defs). symbols.Index must run first; these are
//     tier 2.
//
// New passes are added as separate methods (run*Diagnostics) and concatenated
// into the appropriate tier here. The ordering within a tier matters only for
// tests that assert exact diagnostic order; the LSP client renders them sorted
// by location, and goldens are sorted by (line, column, code, severity).
//
// Current passes:
//
// tree-only (tier 1):
//  0. syntax / missing    (Error/Warning) — diagnostics_syntax.go
//  1. version mismatch    (SeverityError) — diagnostics_version.go
//  2. command version     (SeverityHint)  — diagnostics_version.go
//  3. wrong section       (SeverityHint)  — diagnostics_section.go
//  4. protocol mismatch   (SeverityError) — diagnostics_protocol.go
//     (cross-protocol: an OSPF command inside a BGP router section, etc.;
//     runs after wrong-section because it is the more specific, Error-
//     severity signal, while wrong-section stays Hint)
//
// symbol-dependent (tier 2):
//  5. undefined refs      (SeverityWarning) — diagnostics_refs.go
//  6. duplicate defs      (SeverityWarning) — diagnostics_refs.go

// runTreeDiagnostics runs the passes that need only the parse tree + keyword DB.
// It does NOT consult the symbol table, so it may run before symbols.Index.
func (f *CiscoIOSFeature) runTreeDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	diags := make([]protocol.Diagnostic, 0)
	diags = append(diags, f.runSyntaxDiagnostics(doc, tree)...)
	diags = append(diags, f.runVersionMismatchDiagnostics(doc, tree)...)
	diags = append(diags, f.runCommandVersionDiagnostics(doc, tree)...)
	diags = append(diags, f.runWrongSectionDiagnostics(doc, tree)...)
	diags = append(diags, f.runProtocolMismatchDiagnostics(doc, tree)...)
	return diags
}

// runRefDiagnostics runs the passes that depend on the symbol table
// (undefined-refs, duplicate-defs). symbols.Index must have run first.
func (f *CiscoIOSFeature) runRefDiagnostics(doc *document.Document) []protocol.Diagnostic {
	diags := make([]protocol.Diagnostic, 0)
	diags = append(diags, f.runUndefinedReferenceDiagnostics(doc)...)
	diags = append(diags, f.runDuplicateDefinitionDiagnostics(doc)...)
	return diags
}

// runDiagnostics returns the full set (tree + ref) in one shot. The live LSP
// path (DidOpen/DidChange) instead publishes runTreeDiagnostics as tier 1
// before symbols.Index, then runRefDiagnostics as tier 2. The set is identical
// to this one-shot result; only the internal ordering differs (ref diagnostics
// move to the tail), which no consumer depends on (goldens sort; inline tests
// search by code). Kept as a convenience for one-shot callers (perf harness).
func (f *CiscoIOSFeature) runDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	diags := f.runTreeDiagnostics(doc, tree)
	diags = append(diags, f.runRefDiagnostics(doc)...)
	return diags
}
