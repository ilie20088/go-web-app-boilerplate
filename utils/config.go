package utils

import (
	goflag "flag"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	ConfigManager *viper.Viper = viper.New()
	configpath *string = pflag.String("configpath", "./config", "Path to 'config.yaml' file")
)

func SetDefaults() {
	ConfigManager.SetDefault("host", "localhost")
	ConfigManager.SetDefault("port", "8000")

	ConfigManager.SetDefault("db.host", "localhost")
	ConfigManager.SetDefault("db.name", "dbname")
	ConfigManager.SetDefault("db.user", "dbuser")
	ConfigManager.SetDefault("db.password", "password")

	ConfigManager.SetDefault("log.level", "warn")
	ConfigManager.SetDefault("log.output", []string{"app.log"})
	ConfigManager.SetDefault("log.caller", false)
	ConfigManager.SetDefault("log.stacktrace", true)
}

func ReadFromFile(filename, cfgpath string) error {
	ConfigManager.SetConfigName(filename)
	ConfigManager.AddConfigPath(cfgpath)
	return ConfigManager.ReadInConfig()
}

func GetAddr() string {
	return ConfigManager.GetString("host") + ":" + ConfigManager.GetString("port")
}

func InitConfig() {
	SetDefaults()

	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()

	err := ReadFromFile("config", *configpath)
	if err != nil {
		log.Println(err)
	}
}
