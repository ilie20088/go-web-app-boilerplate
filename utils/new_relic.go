package utils

import (
	"github.com/newrelic/go-agent"
)

// InitNewRelic initializes NewRelic application
func InitNewRelic() (newrelic.Application, error) {
	nrAppName := GetNewRelicAppName()
	nrLicense := GetNewReliceLicense()

	if nrLicense == "" {
		return nil, nil
	}

	Logger.Info("Connecting to New Relic with application name " + nrAppName)
	config := newrelic.NewConfig(nrAppName, nrLicense)

	return newrelic.NewApplication(config)
}
