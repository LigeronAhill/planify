package main

import (
	"log"
	"log/slog"

	"github.com/LigeronAhill/planify/internal/config"
	"github.com/LigeronAhill/planify/internal/telemetry"
)

func main() {
	telemetry.InitLogger()
	configuration, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	token := configuration.GetString("TELEGRAM_BOT_TOKEN")
	slog.Info("Файл конфигурации прочитан", slog.String("token", token))
}
