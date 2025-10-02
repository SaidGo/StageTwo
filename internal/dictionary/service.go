//go:build disable_extras

package dictionary

import (
	"sync"

	"example.com/local/Go2part/dto"
	"example.com/local/Go2part/internal/app"
)

type Service struct {
	*app.BaseService

	lock         sync.RWMutex
	usersByEmail map[string]dto.UserDTO
}
