//go:build disable_extras

package olegalentity

// Минимальная заглушка, чтобы удовлетворить ссылки из пакета web.
type LegalEntityHandler struct{}

// (если web/router.go ожидает функции регистрации, она должна быть в другом файле этого пакета;
// здесь оставляем только тип, которого не хватало)
