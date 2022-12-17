package namespace

import (
	"context"

	"github.com/machado-br/k8s-api/models"
)

type Adapter interface {
	Retrieve(ctx context.Context, name string) (models.Namespace, error)
	Create(ctx context.Context, namespaceModel models.Namespace) error
	RetrieveAll(ctx context.Context) ([]models.Namespace, error)
	Delete(ctx context.Context, name string) error
}
