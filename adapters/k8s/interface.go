package k8s

import (
	"context"

	"github.com/machado-br/k8s-api/models"
	v1 "k8s.io/api/core/v1"
)

type Adapter interface {
	RetrieveNamespaces(ctx context.Context) (*v1.NamespaceList, error)
	CreateNamespace(ctx context.Context, namespace models.Namespace) error
}
