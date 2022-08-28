package helm

import "github.com/machado-br/k8s-api/adapters/models"

type adapter struct{
}

type Adapter interface{
	ListReleases() []models.Release
}

func NewAdapter(
) (adapter, error) {
	return adapter{
	}, nil
}
func (a adapter) ListReleases() []models.Release{
	return []models.Release{}
}
