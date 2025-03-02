package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/LigeronAhill/planify/internal/config"
	"github.com/LigeronAhill/planify/internal/storage"
	"github.com/LigeronAhill/planify/internal/telemetry"
)

func main() {
	ctx := context.Background()
	telemetry.InitLogger()
	configuration, err := config.Init("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	token := configuration.GetString("TELEGRAM_BOT_TOKEN")
	slog.Info("Файл конфигурации прочитан", slog.String("token", token))
	repository, err := storage.New(ctx, "storage.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := repository.Close(); err != nil {
			panic(err)
		}
	}()
}
