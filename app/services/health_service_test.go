package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHealthMessage(t *testing.T) {
	expectedMsg := "UP"
	healthService := NewHealthService()

	actualMsg := healthService.GetHealthMessage()

	assert.Equal(t, actualMsg, expectedMsg)
}
