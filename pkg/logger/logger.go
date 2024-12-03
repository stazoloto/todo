package logger

import (
	"context"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "TODO-APP", log.LstdFlags),
	}
}

type contextKey string

const loggerKey contextKey = "logger"

func (l *Logger) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

func FromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(loggerKey).(*Logger); ok {
		return logger
	}
	return NewLogger()
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.Printf("INFO: %s", msg)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	l.Printf("ERROR: %s", msg)
}
