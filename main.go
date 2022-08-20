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

	log.Printf("result: %v\n", result)

    ca, err := infra.DecodeString(infra.StringValue(result.Cluster.CertificateAuthority.Data))
    if err != nil {
		log.Fatalf("Failed while decoding certificate: %v", err)
    }

	cluster := models.Cluster{
		Arn: infra.StringValue(result.Cluster.Arn),
		Name: infra.StringValue(result.Cluster.Name),
		Endpoint: infra.StringValue(result.Cluster.Endpoint),
		Certificate: ca,
	}

	k8sAdapter, err := k8sAdapter.NewAdapter(cluster)
	if err != nil {
		log.Fatalf("Failed while creating k8s adapter: %v", err)
	}

	secretList, err := k8sAdapter.RetrieveSecret("default")
	if err != nil {
		log.Fatalf("Failed while retrieving k8s secret: %v", err)
	}

	secret := secretList.Items[0]

	err = k8sAdapter.WriteToFile(secret.Data["ca.crt"])
	if err != nil {
		log.Fatalf("Failed while writing kubeconfig file: %v", err)
	}
}
