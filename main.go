package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/dgethings/chunter/lsp"
	"github.com/dgethings/chunter/parse"
	"github.com/dgethings/chunter/rpc"
)

var version = "dev"

func main() {
	logger := getLogger("/Users/dgethings/git/chunter/log.log")
	logger.Println("at the ready")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := parse.NewState()
	defer state.Close()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("error processing data from client: %s", err)
		}
		handleMessage(logger, writer, state, method, contents)
	}

	if err := scanner.Err(); err != nil {
		logger.Printf("scanner error: %v", err)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state parse.State, method string, contents []byte) {
	logger.Printf("got message: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("initialize: %s", err)
			return
		}

		logger.Printf("connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID, version)
		writeResponse(writer, msg)
		logger.Println("sent reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocument
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/DidOpen: %s", err)
			return
		}

		logger.Printf("opened: %s", request.Params.TextDocument.URI)
		state.SetDocument(request.Params.TextDocument.URI, []byte(request.Params.TextDocument.Text))
	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}

		for _, change := range request.Params.ContentChanges {
			diagnostics := state.UpdateDocument(request.Params.TextDocument.URI, []byte(change.Text), logger)
			writeResponse(writer, lsp.PublishDiagnosticsNotification{
				Notification: lsp.Notification{
					RPC:    "2.0",
					Method: "textDocument/publishDiagnostics",
				},
				Params: lsp.PublishDiagnosticsParams{
					URI:         request.Params.TextDocument.URI,
					Diagnostics: diagnostics,
				},
			})
		}
		logger.Printf("changed: %s", request.Params.TextDocument.URI)
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}

		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		writeResponse(writer, response)
	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}

		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		writeResponse(writer, response)
	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/complation: %s", err)
			return
		}

		response := state.Completion(request.ID, request.Params.TextDocument.URI)

		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filenamme string) *log.Logger {
	logfile, err := os.OpenFile(filenamme, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic("Damn, creating file logger no beuno")
	}

	return log.New(logfile, "[chunter]", log.Ldate|log.Ltime|log.Lshortfile)
}
