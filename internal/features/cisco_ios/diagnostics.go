package cisco_ios

import (
	"github.com/dgethings/chunter/internal/ast"
	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func (f *CiscoIOSFeature) runDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}
	if tree == nil {
		return diagnostics
	}
	root := tree.RootNode()
	if root == nil {
		return diagnostics
	}

	runVerNode := root.ChildByFieldName("running_version")
	var cfgVerNode *sitter.Node
	if ss := ast.NamedChildByKind(root, "service_section"); ss != nil {
		cfgVerNode = ss.ChildByFieldName("configured_version")
	}

	var runVer, cfgVer string
	if runVerNode != nil {
		runVer = string(doc.Content[runVerNode.StartByte():runVerNode.EndByte()])
	}
	if cfgVerNode != nil {
		cfgVer = string(doc.Content[cfgVerNode.StartByte():cfgVerNode.EndByte()])
	}
	if runVerNode != nil && cfgVerNode != nil && runVer != cfgVer {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:    protocol.LineRange(runVerNode.StartPosition().Row, runVerNode.StartPosition().Column, runVerNode.EndPosition().Column),
			Severity: 1,
			Source:   "chunter",
			Message:  "running version and configured version mismatch",
		})
	}

	return diagnostics
}
