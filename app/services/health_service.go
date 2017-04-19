package services

type HealthService interface {
	GetHealthMessage() string
}

type HealthServiceImpl struct{}

func (h HealthServiceImpl) GetHealthMessage() string {
	return "UP"
}
