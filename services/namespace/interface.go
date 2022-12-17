package namespace

import (
	"context"

	"github.com/machado-br/k8s-api/models"
)

type Service interface {
	RetrieveAll(ctx context.Context) ([]models.Namespace, error)
	Retrieve(ctx context.Context, name string) (models.Namespace, error)
	CreateNamespace(ctx context.Context, namespace models.Namespace) error
	DeleteNamespace(ctx context.Context, name string) error
}
