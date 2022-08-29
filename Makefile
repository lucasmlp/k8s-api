include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

list-releases:
	@ echo
	@ echo "Listing helm releases on cluster ..."
	@ echo
	@ rm -rf ./config/kube
	@ go run ./cmd/main.go