package namespace

import "k8s.io/client-go/kubernetes"

type adapter struct {
	ClientSet *kubernetes.Clientset
}

func NewAdapter(clientSet *kubernetes.Clientset) Adapter {
	return adapter{
		ClientSet: clientSet,
	}
}
