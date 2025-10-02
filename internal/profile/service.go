//go:build disable_extras

package profile

import "example.com/local/Go2part/internal/app"

type Service struct {
	*app.BaseService
	repo *Repository // Repository объявлен в internal/profile/repository.go и имеет поле rds
}
