package protocol

// ReferenceParams is the parameter object for textDocument/references.
// Embeds TextDocumentPositionParams (textDocument + position) and adds the
// ReferenceContext flag controlling whether the declaration site is also
// returned.
//
// LSP spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_references
type ReferenceParams struct {
	TextDocumentPositionParams
	Context ReferenceContext `json:"context"`
}

// ReferenceContext carries the includeDeclaration flag for a references
// request.
type ReferenceContext struct {
	IncludeDeclaration bool `json:"includeDeclaration"`
}
