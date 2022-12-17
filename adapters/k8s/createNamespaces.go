package k8s

import (
	"context"

	"github.com/machado-br/k8s-api/models"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a adapter) CreateNamespace(ctx context.Context, namespaceModel models.Namespace) error {
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceModel.Name,
		},
	}

	_, err := a.ClientSet.CoreV1().Namespaces().Create(ctx, namespace, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
