package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a adapter) RetrieveNamespaces(ctx context.Context) (*v1.NamespaceList, error) {
	namespaceList, err := a.ClientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return namespaceList, nil
}
