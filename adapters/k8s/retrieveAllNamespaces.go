package k8s

import (
	"context"

	"github.com/machado-br/k8s-api/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a adapter) RetrieveAllNamespaces(ctx context.Context) ([]models.Namespace, error) {
	namespaceList, err := a.ClientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var result []models.Namespace
	for _, namespace := range namespaceList.Items {
		result = append(result, models.Namespace{
			Name:   namespace.Name,
			Status: namespace.Status.Phase,
		})
	}

	return result, nil
}
