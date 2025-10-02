//go:build !disable_extras

package task

// Минимальный alias, достаточный для компиляции ORM-полей.
// Если позже потребуется полноценная работа с БД, заменим на gorm.io/datatypes.JSON.
type JSONB []byte
