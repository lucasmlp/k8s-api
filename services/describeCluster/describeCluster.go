package describeCluster

import (
	"log"

	"github.com/machado-br/k8s-api/adapters/aws"
	"github.com/machado-br/k8s-api/adapters/models"
	"github.com/machado-br/k8s-api/infra"
)

type service struct {
	cloudProviderAdapter aws.Adapter
}

type Service interface{
	Run() (models.Cluster, error)
}

func NewService(
	cloudProviderAdapter aws.Adapter,
) (service, error){
	return service{
		cloudProviderAdapter: cloudProviderAdapter,
	}, nil
}

func (s service) Run() (models.Cluster, error){
	result, err := s.cloudProviderAdapter.DescribeCluster()
	if err != nil {
		log.Fatalf("Failed while calling DescribeCluster: %v", err)
	}

    ca, err := infra.DecodeString(infra.StringValue(result.Cluster.CertificateAuthority.Data))
    if err != nil {
		log.Fatalf("Failed while decoding certificate: %v", err)
    }

	return models.Cluster{
		Arn: infra.StringValue(result.Cluster.Arn),
		Name: infra.StringValue(result.Cluster.Name),
		Endpoint: infra.StringValue(result.Cluster.Endpoint),
		Certificate: ca,
	}, nil
}