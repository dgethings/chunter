package cisco_ios_jinja2

import (
	"regexp"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

var runningVersionRe = regexp.MustCompile(`^!\s*version\s+(\S+)`)

func (f *CiscoIOSFeature) runDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}
	if tree == nil {
		return diagnostics
	}
	root := tree.RootNode()
	if root == nil {
		return diagnostics
	}

	var cfgVerNode, cfgVerField, runVerNode *sitter.Node
	for i := uint(0); i < root.NamedChildCount(); i++ {
		c := root.NamedChild(i)
		if c == nil {
			continue
		}
		if cfgVerNode == nil && c.Kind() == "version_statement" {
			cfgVerNode = c
			cfgVerField = c.ChildByFieldName("configured_version")
		}
		if runVerNode == nil && c.Kind() == "comment" {
			runVerNode = c
		}
	}

	var runVer, cfgVer string
	if runVerNode != nil {
		runVerText := string(doc.Content[runVerNode.StartByte():runVerNode.EndByte()])
		if m := runningVersionRe.FindStringSubmatch(runVerText); m != nil {
			runVer = m[1]
		}
	}
	if cfgVerField != nil {
		cfgVer = string(doc.Content[cfgVerField.StartByte():cfgVerField.EndByte()])
	}
	if runVer != "" && cfgVer != "" && runVer != cfgVer {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:    protocol.LineRange(cfgVerNode.StartPosition().Row, cfgVerNode.StartPosition().Column, cfgVerNode.EndPosition().Column),
			Severity: 1,
			Source:   "chunter",
			Message:  "running version and configured version mismatch",
		})
	}

	return diagnostics
}
