package logger

import (
	"log/slog"
	"os"

	"github.com/pirotyyy/todo-app-with-ddd-cicd-iac/internal/config"
)

var l *slog.Logger

func InitLogger(confLog *config.Log) {
	opts := &slog.HandlerOptions{
		Level: slogLevel(confLog.Level),
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)

	l = slog.New(handler)
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

func Info(msg string, keysAndValues ...any) {
	l.Info(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...any) {
	l.Debug(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...any) {
	l.Warn(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...any) {
	l.Error(msg, keysAndValues...)
}
