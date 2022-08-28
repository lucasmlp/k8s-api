package listReleases

import (
	"github.com/machado-br/k8s-api/adapters/helm"
	"github.com/machado-br/k8s-api/adapters/models"
)

type service struct{
	helmAdapter helm.Adapter
}

type Service interface{
	Run() ([]models.Release, error)
}

func NewService(
	helmAdapter helm.Adapter,
) (service, error) {
	return service{
		helmAdapter: helmAdapter,
	}, nil
}

func (s service) Run() ([]models.Release, error) {
	return []models.Release{}, nil
}