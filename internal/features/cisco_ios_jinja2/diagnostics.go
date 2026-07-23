package cisco_ios_jinja2

import (
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// runDiagnostics is the dispatcher over every diagnostic pass. New passes
// are added as separate methods (run*Diagnostic) and concatenated here. The
// ordering matters only for tests that assert exact diagnostic order; the
// LSP client renders them sorted by location.
//
// Current passes, in order:
//  1. version mismatch  (SeverityError)   — diagnostics_version.go
//  2. undefined refs    (SeverityWarning) — diagnostics_refs.go
//  3. duplicate defs    (SeverityWarning) — diagnostics_refs.go
//  4. wrong section     (SeverityHint)    — diagnostics_section.go
func (f *CiscoIOSFeature) runDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	var diags []protocol.Diagnostic
	diags = append(diags, f.runVersionMismatchDiagnostics(doc, tree)...)
	diags = append(diags, f.runUndefinedReferenceDiagnostics(doc)...)
	diags = append(diags, f.runDuplicateDefinitionDiagnostics(doc)...)
	diags = append(diags, f.runWrongSectionDiagnostics(doc, tree)...)
	return diags
}
