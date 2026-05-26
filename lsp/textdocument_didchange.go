package lsp

type TextDocumentDidChangeNotification struct {
	Notification
	Params DidChangeTextDocumentParmas `json:"params"`
}

type DidChangeTextDocumentParmas struct {
	TextDocument   VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentchanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
