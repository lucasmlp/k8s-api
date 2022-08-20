include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

kube-file:
	@ echo
	@ echo "Generating kubeconfig file.."
	@ echo
	@ rm -rf kubeconfig
	@ go run ./cmd/main.go