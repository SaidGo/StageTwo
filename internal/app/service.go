//go:build disable_extras

package app

import "context"

// Service — общий контракт на будущее (минимален для компиляции).
type Service interface {
	Health(ctx context.Context) error
}

// BaseService — простая заглушка.
type BaseService struct{}

func (s *BaseService) Health(ctx context.Context) error { return nil }
