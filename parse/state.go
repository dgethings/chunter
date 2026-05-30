package parse

import (
	"fmt"
	"log"

	"github.com/dgethings/chunter/lsp"
	ts_ci "github.com/dgethings/chunter/tree-sitter-cisco_ios/bindings/go"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type State struct {
	Documents map[string][]byte
	Parser    *ts.Parser
	Tree      *ts.Tree
}

func NewState() State {
	return State{
		Documents: map[string][]byte{},
		Parser:    ts.NewParser(),
	}
}

func (s *State) Close() {
	s.Parser.Close()
	s.Tree.Close()
}

func (s *State) SetDocument(uri string, text []byte) {
	s.Documents[uri] = text
	s.Parser.SetLanguage(ts.NewLanguage(ts_ci.Language()))
	s.Tree = s.Parser.Parse([]byte(text), nil)
}

func (s *State) UpdateDocument(uri string, text []byte, logger *log.Logger) []lsp.Diagnostic {
	s.Documents[uri] = text
	s.Tree = s.Parser.Parse([]byte(text), s.Tree)
	diagnostics := []lsp.Diagnostic{}
	if s.Tree == nil {
		logger.Printf("parser returned empty tree for: %s", text)
		return diagnostics
	}
	root := s.Tree.RootNode()
	if root == nil {
		logger.Println("tree has no root node")
		return diagnostics
	}
	runVerNode := root.ChildByFieldName("running_version")
	var runVer string
	if runVerNode != nil {
		runVer = string(text[runVerNode.StartByte():runVerNode.EndByte()])
	}
	cfgVerNode := root.ChildByFieldName("configured_version")
	var cfgVer string
	if cfgVerNode != nil {
		cfgVer = string(text[cfgVerNode.StartByte():cfgVerNode.EndByte()])
	}
	if runVer != cfgVer {
		logger.Printf("version mismatch. running: %s configured: %s", runVer, cfgVer)

		diagnostics = append(diagnostics, lsp.Diagnostic{
			Range:    LineRange(runVerNode.StartPosition().Row, runVerNode.StartPosition().Column, runVerNode.EndPosition().Column),
			Severity: 1,
			Source:   "chunter",
			Message:  "running version and configured version mismatch",
		})
	}

	// for row, line := range strings.Split(text, "\n") {
	// 	if strings.Contains(line, "cisco") {
	// 		idx := strings.Index(line, "cisco")
	// 		diagnostics = append(diagnostics, lsp.Diagnostic{
	// 			Range:    LineRange(uint(row), uint(idx), uint(idx+len("cisco"))),
	// 			Severity: 1,
	// 			Source:   "chunter",
	// 			Message:  "watch for the 800lb gorilla",
	// 		})
	// 	}
	// }
	return diagnostics
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// TODO: real hover response

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File %s, Characters: %d", uri, len(document)),
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// TODO: real definition response

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) Completion(id int, uri string) lsp.CompletionResponse {

	completions := []lsp.CompletionItem{
		{
			Label:         "NAF",
			Detail:        "Not going this year :(",
			Documentation: "First one in EU that I'm missing. V sad",
		},
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: completions,
	}
}

func LineRange(line, start, end uint) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
