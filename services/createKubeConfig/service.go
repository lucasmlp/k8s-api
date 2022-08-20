package createKubeConfig

import (
	"log"

	"github.com/machado-br/k8s-api/adapters/k8s"
)

type service struct {
	clusterName string
	clusterRegion string
	k8sAdapter k8s.Adapter
}

type Service interface{
	Run() error
}

func NewService(
	clusterName string,
	clusterRegion string,
	k8sAdapter k8s.Adapter,
) (service, error){
	return service{
		clusterName: clusterName,
		clusterRegion: clusterRegion,
		k8sAdapter: k8sAdapter,
	}, nil
}

func (s service) Run() error{

	secretList, err := s.k8sAdapter.RetrieveSecret("default")
	if err != nil {
		log.Fatalf("Failed while retrieving k8s secret: %v", err)
	}

	if len(secretList.Items) == 0 {
		log.Panicln("Secret list is empty")
	}

	secret := secretList.Items[0]

	err = s.k8sAdapter.WriteToFile(secret.Data["ca.crt"])
	if err != nil {
		log.Fatalf("Failed while writing kubeconfig file: %v", err)
	}

	return nil
}