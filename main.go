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
	secretList, err := clientset.CoreV1().Secrets("dev").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	secret := secretList.Items[1]

	clustersList := map[string]*api.Cluster{
		*result.Cluster.Arn: {
			Server:                   *result.Cluster.Endpoint,
			CertificateAuthorityData: secret.Data["ca.crt"],
		},
	}

	contextList := map[string]*api.Context{
		*result.Cluster.Arn: {
			Cluster:  *result.Cluster.Arn,
			AuthInfo: *result.Cluster.Arn,
		},
	}

	execEnvList := []api.ExecEnvVar{
		{
			Name:  "AWS_PROFILE",
			Value: "ihm",
		},
	}

	exec := api.ExecConfig{
		Command:    "aws",
		Args:       []string{"eks", "get-token", "--region", "us-west-2", "--cluster-name", "zetta-non-prod"},
		Env:        execEnvList,
		APIVersion: "client.authentication.k8s.io/v1beta1",
	}

	authInfoList := map[string]*api.AuthInfo{
		*result.Cluster.Arn: {
			Exec: &exec,
		},
	}

	clientConfig := api.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clustersList,
		Contexts:       contextList,
		AuthInfos:      authInfoList,
		CurrentContext: *result.Cluster.Arn,
	}

	err = clientcmd.WriteToFile(clientConfig, "./kubeconfig")
	if err != nil {
		log.Fatalln(err)
	}
}
