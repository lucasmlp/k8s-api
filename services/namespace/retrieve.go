package namespace

import (
	"context"

	"github.com/machado-br/k8s-api/models"
)

func (s *service) Retrieve(ctx context.Context, name string) (models.Namespace, error) {
	result, err := s.Adapter.Retrieve(ctx, name)
	if err != nil {
		return models.Namespace{}, err
	}

	return result, nil
}
