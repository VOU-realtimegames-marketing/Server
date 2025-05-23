package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables
type Config struct {
	Environment              string `mapstructure:"ENVIRONMENT"`
	HTTPServerAddress        string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress        string `mapstructure:"GRPC_SERVER_ADDRESS"`
	AuthServerAddress        string `mapstructure:"AUTH_SERVER_ADDRESS"`
	CounterpartServerAddress string `mapstructure:"COUNTERPART_SERVER_ADDRESS"`
	EventServerAddress       string `mapstructure:"EVENT_SERVER_ADDRESS"`
	GameServerAddress        string `mapstructure:"GAME_SERVER_ADDRESS"`
	UserServerAddress        string `mapstructure:"USER_SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
