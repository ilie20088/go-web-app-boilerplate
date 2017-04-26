package utils

import (
	goflag "flag"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

var (
	configManager = viper.New()
	configpath    = pflag.String("configpath", "./config", "Path to 'config.yaml' file")
)

// SetDefaults sets default values for all configuration parameters
func SetDefaults() {
	configManager.SetDefault("host", "localhost")
	configManager.SetDefault("port", "8000")

	configManager.SetDefault("db.host", "localhost")
	configManager.SetDefault("db.name", "dbname")
	configManager.SetDefault("db.user", "dbuser")
	configManager.SetDefault("db.password", "password")
	configManager.SetDefault("db.max-open-connections", 250)
	configManager.SetDefault("db.max-idle-connections", 100)
	configManager.SetDefault("db.conn-max-lifetime", 180*time.Second)

	configManager.SetDefault("cache.host", "localhost")
	configManager.SetDefault("cache.port", "6379")

	configManager.SetDefault("log.level", "warn")
	configManager.SetDefault("log.output", []string{"app.log"})
	configManager.SetDefault("log.caller", false)
	configManager.SetDefault("log.stacktrace", true)

	configManager.SetDefault("newrelic.app-name", "")
	configManager.SetDefault("newrelic.license", "")
}

func readFromFile(filename, cfgpath string) error {
	configManager.SetConfigName(filename)
	configManager.AddConfigPath(cfgpath)
	return configManager.ReadInConfig()
}

// GetAddr returns address service should run on (e.g. localhost:8000)
func GetAddr() string {
	return configManager.GetString("host") + ":" + configManager.GetString("port")
}

// GetDBHost returns host name of database instance
func GetDBHost() string {
	return configManager.GetString("db.host")
}

// GetDBName returns database name
func GetDBName() string {
	return configManager.GetString("db.name")
}

// GetDBUser returns database user
func GetDBUser() string {
	return configManager.GetString("db.user")
}

// GetDBpass returns database user's password
func GetDBpass() string {
	return configManager.GetString("db.pass")
}

// GetMaxOpenConnections returns maximum number of open connections to database
func GetMaxOpenConnections() int {
	return configManager.GetInt("db.max-open-connections")
}

// GetMaxIdleConnections returns maximum number of idle connections to database
func GetMaxIdleConnections() int {
	return configManager.GetInt("db.max-idle-connections")
}

// GetConnectionMaxLifetime returns maximum lifetime of connection
func GetConnectionMaxLifetime() time.Duration {
	return configManager.GetDuration("db.conn-max-lifetime")
}

// GetCacheBDAddr returns address of cache instance
func GetCacheBDAddr() string {
	return configManager.GetString("cache.host") + ":" + configManager.GetString("cache.port")
}

// GetLogLevel returns current logging level
func GetLogLevel() string {
	return configManager.GetString("log.level")
}

// GetLogOutput returns files where logs should go
func GetLogOutput() []string {
	return configManager.GetStringSlice("log.output")
}

// ShouldLogCaller returns whether we should caller
func ShouldLogCaller() bool {
	return configManager.GetBool("log.caller")
}

// ShouldLogStacktrace returns whether we should log stacktrace
func ShouldLogStacktrace() bool {
	return configManager.GetBool("log.stacktrace")
}

// GetNewRelicAppName returns NewRelic application name
func GetNewRelicAppName() string {
	return configManager.GetString("newrelic.app-name")
}

// GetNewReliceLicense returns NewRelic license
func GetNewReliceLicense() string {
	return configManager.GetString("newrelic.license")
}

// InitConfig initializes configs
func InitConfig() {
	SetDefaults()

	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()

	err := readFromFile("config", *configpath)
	if err != nil {
		log.Println(err)
	}
}
