aws ecr create-repository \
    --repository-name k8s-api \
    --image-scanning-configuration scanOnPush=true \
    --region us-west-2