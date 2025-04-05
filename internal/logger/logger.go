package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/pirotyyy/todo-app-with-ddd-cicd-iac/internal/config"
)

const (
	skipCount = 3
)

type Handler struct {
	slog.Handler
}

func NewHandler(confLog *config.Log) *Handler {
	opts := &slog.HandlerOptions{
		Level: slogLevel(confLog.Level),
	}

	return &Handler{
		Handler: slog.NewJSONHandler(os.Stdout, opts),
	}
}

func slogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func Info(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelInfo, msg, args...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelDebug, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelWarn, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelError, msg, args...)
}

func log(ctx context.Context, level slog.Level, msg string, keysAndValues ...any) {
	logger := slog.Default()
	if !logger.Enabled(ctx, level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(skipCount, pcs[:])

	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(keysAndValues...)

	_ = logger.Handler().Handle(ctx, r)
}
