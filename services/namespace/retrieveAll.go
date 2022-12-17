package namespace

import (
	"context"

	"github.com/machado-br/k8s-api/models"
)

func (s *service) RetrieveAll(ctx context.Context) ([]models.Namespace, error) {
	namespaceList, err := s.Adapter.RetrieveAll(ctx)
	if err != nil {
		return nil, err
	}
	return namespaceList, nil
}
