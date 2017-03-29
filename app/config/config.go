package config

import (
	"github.com/spf13/viper"
)

var (
	Manager *viper.Viper = viper.New()
)

func SetDefaults() {
	Manager.SetDefault("host", "localhost")
	Manager.SetDefault("port", "8000")
	Manager.SetDefault("db.host", "localhost")
	Manager.SetDefault("db.name", "playheads")
	Manager.SetDefault("db.user", "playheads")
	Manager.SetDefault("db.pass", "playheads")
}

func ReadFromFile() error {
	Manager.SetConfigName("config")
	Manager.AddConfigPath("config/")
	return Manager.ReadInConfig()
}

func GetAddr() string {
	return Manager.GetString("host") + ":" + Manager.GetString("port")
}
