package main

import (
	"log"

	awsAdapter "github.com/machado-br/k8s-api/adapters/aws"
	k8sAdapter "github.com/machado-br/k8s-api/adapters/k8s"
	"github.com/machado-br/k8s-api/adapters/models"
	"github.com/machado-br/k8s-api/infra"
)

func main() {
    name := "e-commerce"
    region := "us-west-2"

	cloudProviderAdapter, err := awsAdapter.NewAdapter(region, name)
	if err != nil {
		log.Fatalf("Failed while creating cloud provider adapter: %v", err)
	}

	result, err := cloudProviderAdapter.DescribeCluster()
	if err != nil {
		log.Fatalf("Failed while calling DescribeCluster: %v", err)
	}

	k8sAdapter, err := k8sAdapter.NewAdapter()
	if err != nil {
		log.Fatalf("Failed while creating k8s adapter: %v", err)
	}

    ca, err := infra.DecodeString(infra.StringValue(result.Cluster.CertificateAuthority.Data))
    if err != nil {
		log.Fatalf("Failed while decoding certificate: %v", err)
    }

	cluster := models.Cluster{
		Name: infra.StringValue(result.Cluster.Name),
		Endpoint: infra.StringValue(result.Cluster.Endpoint),
		Certificate: ca,
	}

    _, err = k8sAdapter.NewClientset(cluster)
    if err != nil {
        log.Fatalf("Error creating clientset: %v", err)
    }

}