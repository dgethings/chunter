package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
	// Params: added in the per request type
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
