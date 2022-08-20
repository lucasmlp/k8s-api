package adapters

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

type adapter struct{
	region string
	clusterName string
	session *session.Session
	eks *eks.EKS
}

type Adapter interface{
	DescribeCluster() (*eks.DescribeClusterOutput, error)
}

func NewAdapter(
	region string,
	clusterName string,
) (adapter, error) {
	sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(region),
    }))

    eksSvc := eks.New(sess)

	return adapter{
		region: region,
		clusterName: clusterName,
		session: sess,
		eks: eksSvc,
	}, nil
}

func (a adapter) DescribeCluster() (*eks.DescribeClusterOutput, error){
	
    input := &eks.DescribeClusterInput{
        Name: aws.String(a.clusterName),
    }

    result, err := a.eks.DescribeCluster(input)
    if err != nil {
        return nil, err
    }

	return result, nil
}