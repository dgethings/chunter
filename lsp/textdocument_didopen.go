package lsp

type DidOpenTextDocument struct {
	Notification
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
