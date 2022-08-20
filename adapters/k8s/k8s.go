package k8s

import (
	"log"

	"github.com/machado-br/k8s-api/adapters/models"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

type adapter struct{
}

type Adapter interface{
	NewClientset(cluster models.Cluster) (*kubernetes.Clientset, error)
}

func NewAdapter(
) (adapter, error) {
	return adapter{
	}, nil
}

func (a adapter) NewClientset(cluster models.Cluster) (*kubernetes.Clientset, error) {
    log.Printf("Cluster name: %+v", cluster.Name)

    gen, err := token.NewGenerator(true, false)
    if err != nil {
        return nil, err
    }

    opts := &token.GetTokenOptions{
        ClusterID: cluster.Name,
    }

    tok, err := gen.GetWithOptions(opts)
    if err != nil {
        return nil, err
    }

    clientset, err := kubernetes.NewForConfig(
        &rest.Config{
            Host:        cluster.Endpoint,
            BearerToken: tok.Token,
            TLSClientConfig: rest.TLSClientConfig{
                CAData: cluster.Certificate,
            },
        },
    )

    if err != nil {
        return nil, err
    }

    return clientset, nil
}