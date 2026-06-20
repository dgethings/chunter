package server

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/handler"
	"github.com/dgethings/chunter/internal/logger"
	"github.com/dgethings/chunter/internal/protocol"
)

type InitializeParams struct {
	ClientInfo *protocol.ClientInfo `json:"clientInfo"`
}

func (s *Server) Initialize(ctx context.Context, params InitializeParams) (protocol.InitializeResult, error) {
	s.setState(stateInitialized)
	logger.FromContext(ctx).Printf("connected to: %s %s", params.ClientInfo.Name, params.ClientInfo.Version)
	return protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync:   1,
			HoverProvider:      true,
			DefinitionProvider: true,
			CompletionProvider: &protocol.CompletionOptions{},
		},
		ServerInfo: protocol.ServerInfo{
			Name:    "chunter",
			Version: s.version,
		},
	}, nil
}

func (s *Server) Initialized(ctx context.Context) error {
	logger.FromContext(ctx).Println("client initialized")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.setState(stateShutDown)
	return nil
}

func (s *Server) Assigner() (jrpc2.Assigner, error) {
	return handler.Map{
		"initialize":              handler.New(s.Initialize),
		"initialized":             handler.New(s.Initialized),
		"shutdown":                handler.New(s.Shutdown),
		"textDocument/didOpen":    handler.New(s.DidOpen),
		"textDocument/didChange":  handler.New(s.DidChange),
		"textDocument/didClose":   handler.New(s.DidClose),
		"textDocument/completion": handler.New(s.Completion),
		"textDocument/hover":      handler.New(s.Hover),
		"textDocument/definition": handler.New(s.Definition),
	}, nil
}
