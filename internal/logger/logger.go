package logger

import (
	"context"
	"log"
	"os"
)

type ctxKey struct{}

var defaultLogger = log.New(os.Stderr, "[chunter] ", log.Ldate|log.Ltime|log.Lshortfile)

func Default() *log.Logger {
	return defaultLogger
}

func With(ctx context.Context, l *log.Logger) context.Context {
	if l == nil {
		return ctx
	}
	return context.WithValue(ctx, ctxKey{}, l)
}

func FromContext(ctx context.Context) *log.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*log.Logger); ok {
		return l
	}
	return defaultLogger
}
