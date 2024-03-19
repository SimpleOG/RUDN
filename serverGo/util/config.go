package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	MigrationUrl      string `mapstructure:"MIGRATION_URL"`
	DBDSource         string `mapstructure:"DB_SOURCE"`
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func InitConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return config, err

}
