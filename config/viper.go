package config

import "github.com/spf13/viper"

func initViper() {
	viper.SetConfigName("application")
	viper.WatchConfig()
}
