package logger

import (
	"log/slog"
	"os"
	"research-data-saver/internal/config"

	"github.com/lmittmann/tint"
)

func Init(cfg *config.Config) *slog.Logger {
	var level slog.Level

	switch cfg.Env {
	case "local":
		level = slog.LevelDebug
	case "dev":
		level = slog.LevelDebug
	case "prod":
		level = slog.LevelInfo
	}

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:     level,
		AddSource: true,
	}))

	return logger
}
