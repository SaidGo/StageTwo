//go:build disable_extras

package activities

import "example.com/local/Go2part/internal/app"

type Service struct {
	*app.BaseService
}

func (s *Service) CreateActivity(a *Activity) error { return nil }
