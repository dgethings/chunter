package logger

import (
	"log/slog"
	"os"
)

func SetLogger() {
	l := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("lsp", "chunter")
	slog.SetDefault(l)
}
