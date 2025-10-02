//go:build disable_extras

package reminders

import "example.com/local/Go2part/internal/app"

// Service — локальная обёртка для базового сервиса приложения.
type Service struct {
	*app.BaseService
}
