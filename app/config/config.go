package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	ConfigManager *viper.Viper = viper.New()
)

func SetDefaults() {
	ConfigManager.SetDefault("host", "localhost")
	ConfigManager.SetDefault("port", "8000")
	ConfigManager.SetDefault("db.host", "localhost")
	ConfigManager.SetDefault("db.name", "playheads")
	ConfigManager.SetDefault("db.user", "playheads")
	ConfigManager.SetDefault("db.pass", "playheads")
}

func ReadFromFile(filename, cfgpath string) error {
	ConfigManager.SetConfigName(filename)
	ConfigManager.AddConfigPath(cfgpath)
	return ConfigManager.ReadInConfig()
}

func GetAddr() string {
	return ConfigManager.GetString("host") + ":" + ConfigManager.GetString("port")
}

func init() {
	SetDefaults()
	err := ReadFromFile("config", "config")
	if err != nil {
		log.Fatal(err)
	}
}
