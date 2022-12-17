package namespace

import (
	"context"

	"github.com/machado-br/k8s-api/models"
)

type Service interface {
	CreateNamespace(ctx context.Context, namespace models.Namespace) error
}
