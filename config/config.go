package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Secret string
}

var _config Config

func init() {
	viper.SetConfigName("config")                           // name of config file (without extension)
	viper.SetConfigType("yaml")                             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/code/forum_project/forum_server") // optionally look for config in the working directory
	err := viper.ReadInConfig()                             // Find and read the config file
	if err != nil {                                         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&_config); err != nil {
		panic(err)
	}
}

func GetConfig() string {
	return _config.Secret
}
