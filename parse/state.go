package parse

import (
	"fmt"
	"log"
	"strings"

	"github.com/dgethings/chunter/lsp"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) SetDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string, logger *log.Logger) []lsp.Diagnostic {
	s.Documents[uri] = text
	diagnostics := []lsp.Diagnostic{}

	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "cisco") {
			idx := strings.Index(line, "cisco")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(uint(row), uint(idx), uint(idx+len("cisco"))),
				Severity: 1,
				Source:   "chunter",
				Message:  "watch for the 800lb gorilla",
			})
		}
	}
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
