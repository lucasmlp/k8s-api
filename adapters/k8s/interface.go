package k8s

import (
	"context"

	"github.com/machado-br/k8s-api/models"
)

type Adapter interface {
	RetrieveAllNamespaces(ctx context.Context) ([]models.Namespace, error)
	CreateNamespace(ctx context.Context, namespace models.Namespace) error
	DeleteNamespace(ctx context.Context, name string) error
}
