package namespace

import (
	"context"

	"github.com/machado-br/k8s-api/models"
)

func (s *service) CreateNamespace(ctx context.Context, namespace models.Namespace) error {
	err := s.Adapter.Create(ctx, namespace)
	if err != nil {
		return err
	}
	return nil
}
