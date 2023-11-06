package util

import (
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBDSource     string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func InitConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	file, err := filepath.Glob("*.env")
	viper.SetConfigType("env")
	viper.SetConfigName(file[0])

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return config, err

}
