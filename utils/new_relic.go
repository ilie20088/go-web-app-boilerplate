package utils

import (
	"github.com/newrelic/go-agent"
)

// InitNewRelic initializes NewRelic application
func InitNewRelic() (newrelic.Application, error) {
	newRelicAppName := GetNewRelicAppName()
	newRelicLicense := GetNewReliceLicense()

	if newRelicLicense == "" {
		return nil, nil
	}

	Logger.Info("Connecting to New Relic with application name " + newRelicAppName)
	config := newrelic.NewConfig(newRelicAppName, newRelicLicense)

	return newrelic.NewApplication(config)
}
