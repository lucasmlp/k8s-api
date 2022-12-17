package namespace

import (
	"github.com/machado-br/k8s-api/adapters/k8s/namespace"
)

type service struct {
	Adapter namespace.Adapter
}

func NewService(
	adapter namespace.Adapter,
) Service {
	return &service{
		Adapter: adapter,
	}
}
