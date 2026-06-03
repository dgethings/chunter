package cisco_ios

import (
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
	var runVer string
	if runVerNode != nil {
		runVer = string(doc.Content[runVerNode.StartByte():runVerNode.EndByte()])
	}
	cfgVerNode := root.ChildByFieldName("configured_version")
	var cfgVer string
	if cfgVerNode != nil {
		cfgVer = string(doc.Content[cfgVerNode.StartByte():cfgVerNode.EndByte()])
	}
	if runVer != cfgVer {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:    protocol.LineRange(runVerNode.StartPosition().Row, runVerNode.StartPosition().Column, runVerNode.EndPosition().Column),
			Severity: 1,
			Source:   "chunter",
			Message:  "running version and configured version mismatch",
		})
	}

	return diagnostics
}
