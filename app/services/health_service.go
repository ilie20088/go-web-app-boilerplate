package services

type HealthService interface {
	GetHealthMessage() string
}

type healthServiceImpl struct{}

func NewHealthService() HealthService {
	return &healthServiceImpl{}
}

func (h healthServiceImpl) GetHealthMessage() string {
	return "UP"
}
