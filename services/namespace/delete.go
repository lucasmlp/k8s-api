package namespace

import (
	"context"
)

func (s *service) DeleteNamespace(ctx context.Context, name string) error {
	err := s.Adapter.Delete(ctx, name)
	if err != nil {
		return err
	}
	return nil
}
