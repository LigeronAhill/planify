package config

import (
	"os"

	"github.com/spf13/viper"
)

func Init(fileName string) (*viper.Viper, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	c := viper.New()
	c.SetConfigFile(fileName)
	c.AddConfigPath(cwd)
	err = c.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return c, nil
}
