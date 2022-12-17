package namespace

import (
	"context"

	"github.com/machado-br/k8s-api/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a adapter) Retrieve(ctx context.Context, name string) (models.Namespace, error) {
	namespace, err := a.ClientSet.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return models.Namespace{}, err
	}

	return models.Namespace{
		Name: namespace.Name,
	}, nil
}
