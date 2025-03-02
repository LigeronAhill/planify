package config

import "github.com/spf13/viper"

func Init() (*viper.Viper, error) {
	c := viper.New()
	c.SetConfigFile("config.toml")
	c.AddConfigPath(".")
	err := c.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return c, nil
}
