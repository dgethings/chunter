package cisco_ios_jinja2

import (
	"regexp"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

var runningVersionRe = regexp.MustCompile(`^!\s*version\s+(\S+)`)

// runVersionMismatchDiagnostics reports when the running version (recorded
// in a `! version X` comment, as emitted by `show run`) differs from the
// configured `version Y` statement. The two are normally kept in sync by
// IOS; a mismatch usually means the file was hand-edited or the running
// image was downgraded without rewriting the config.
//
// All top-level named children are scanned: the running version is taken
// from the first comment whose text matches runningVersionRe (first wins,
// matching show-run emit order); the configured version is taken from the
// first version_statement. If only one is present, no diagnostic is
// emitted (that is not an error condition).
func (f *CiscoIOSFeature) runVersionMismatchDiagnostics(doc *document.Document, tree *sitter.Tree) []protocol.Diagnostic {
	var diags []protocol.Diagnostic
	if tree == nil {
		return diags
	}
	root := tree.RootNode()
	if root == nil {
		return diags
	}

	var cfgVerNode, cfgVerField *sitter.Node
	var runVer string
	for i := uint(0); i < root.NamedChildCount(); i++ {
		c := root.NamedChild(i)
		if c == nil {
			continue
		}
		if cfgVerNode == nil && c.Kind() == "version_statement" {
			cfgVerNode = c
			cfgVerField = c.ChildByFieldName("configured_version")
		}
		if runVer == "" && c.Kind() == "comment" {
			text := string(doc.Content[c.StartByte():c.EndByte()])
			if m := runningVersionRe.FindStringSubmatch(text); m != nil {
				runVer = m[1]
			}
		}
	}

	var cfgVer string
	if cfgVerField != nil {
		cfgVer = string(doc.Content[cfgVerField.StartByte():cfgVerField.EndByte()])
	}
	if runVer != "" && cfgVer != "" && runVer != cfgVer {
		diags = append(diags, protocol.Diagnostic{
			Range:    protocol.LineRange(cfgVerNode.StartPosition().Row, cfgVerNode.StartPosition().Column, cfgVerNode.EndPosition().Column),
			Severity: protocol.SeverityError,
			Source:   "chunter",
			Code:     "version-mismatch",
			Message:  "running version and configured version mismatch",
		})
	}

	return diags
}
