package features

import (
	"context"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/protocol"
)

type Feature interface {
	LanguageID() string

	DidOpen(ctx context.Context, doc *document.Document) ([]protocol.Diagnostic, error)
	DidChange(ctx context.Context, doc *document.Document) ([]protocol.Diagnostic, error)
	DidClose(ctx context.Context, doc *document.Document) error

	Completion(ctx context.Context, doc *document.Document, pos protocol.Position) ([]protocol.CompletionItem, error)
	Hover(ctx context.Context, doc *document.Document, pos protocol.Position) (*protocol.HoverResult, error)
	Definition(ctx context.Context, doc *document.Document, pos protocol.Position) ([]protocol.Location, error)
}
