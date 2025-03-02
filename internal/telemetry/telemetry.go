package telemetry

import (
	"log/slog"
	"os"

	prettylogger "github.com/jacute/prettylogger"
)

func InitLogger() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(prettylogger.NewColoredHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	slog.Info("Сообщения", slog.String("INFO", "будут выведены"))
	slog.Debug("Сообщения", slog.String("DEBUG", "будут выведены"))
	slog.Warn("Сообщения", slog.String("WARN", "будут выведены"))
	slog.Error("Сообщения", slog.String("ERROR", "будут выведены"))
}
