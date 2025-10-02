package dto

import (
	"errors"
	"fmt"
)

// Базовая причина "не найдено"
var NotFoundErr = errors.New("not found")

// NotFoundErrMsg — совместимость со старыми вызовами вида dto.NotFoundErr("...").
func NotFoundErrMsg(msg string) error {
	return fmt.Errorf("%w: %s", NotFoundErr, msg)
}

// NotFoundErrf — форматированный вариант ("not found: ...").
func NotFoundErrf(format string, args ...any) error {
	return fmt.Errorf("%w: "+format, append([]any{NotFoundErr}, args...)...)
}
