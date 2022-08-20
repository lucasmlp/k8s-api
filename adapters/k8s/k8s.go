package k8s

import (
	"log"

	"github.com/machado-br/k8s-api/adapters/models"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

type adapter struct{
    cluster models.Cluster
    clientSet *kubernetes.Clientset
}

type Adapter interface{
    RetrieveSecret(namespace string ) (*v1.SecretList, error)
    WriteToFile(certificate []byte) error 
}

func NewAdapter(
    cluster models.Cluster,
) (adapter, error) {
    clientSet, err := newClientset(cluster)
    if err != nil {
        return adapter{}, err
    }

	return adapter{
        cluster: cluster,
        clientSet: clientSet,
	}, nil
}

func newClientset(cluster models.Cluster) (*kubernetes.Clientset, error) {
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

func (a adapter) RetrieveSecret(namespace string ) (*v1.SecretList, error){
    secretList, err := a.clientSet.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

    return secretList, nil
}

func (a adapter) WriteToFile(certificate []byte) error {
    clustersList := map[string]*api.Cluster{
		a.cluster.Arn: {
			Server:                   a.cluster.Endpoint,
			CertificateAuthorityData: certificate,
		},
	}

	contextList := map[string]*api.Context{
		a.cluster.Arn: {
			Cluster:  a.cluster.Arn,
			AuthInfo: a.cluster.Arn,
		},
	}

	exec := api.ExecConfig{
		Command:    "aws",
		Args:       []string{"eks", "get-token", "--region", "us-west-2", "--cluster-name", a.cluster.Name},
		APIVersion: "client.authentication.k8s.io/v1beta1",
	}

	authInfoList := map[string]*api.AuthInfo{
		a.cluster.Arn: {
			Exec: &exec,
		},
	}

	clientConfig := api.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clustersList,
		Contexts:       contextList,
		AuthInfos:      authInfoList,
		CurrentContext: a.cluster.Arn,
	}

	err := clientcmd.WriteToFile(clientConfig, "./kubeconfig")
	if err != nil {
		return err
	}

    return nil
}