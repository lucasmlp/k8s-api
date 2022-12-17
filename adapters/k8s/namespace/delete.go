package namespace

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a adapter) Delete(ctx context.Context, name string) error {
	err := a.ClientSet.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
