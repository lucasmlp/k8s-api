package namespace

import "github.com/machado-br/k8s-api/adapters/k8s"

type service struct {
	Adapter k8s.Adapter
}

func NewService(adapter k8s.Adapter) Service {
	return &service{
		Adapter: adapter,
	}
}
