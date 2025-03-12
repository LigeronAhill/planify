package config

import (
	"log/slog"
	"os/exec"
	"path"
	"strings"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/spf13/viper"
)

func Init(fileName string) (*viper.Viper, error) {
	op := "чтение файла конфигурации"
	slog.Info(op)
	cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	projectPath := strings.TrimSpace(string(cmdOut))
	filePath := path.Join(projectPath, fileName)

	c := viper.New()
	c.SetConfigFile(filePath)

	err = c.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return c, nil
}
