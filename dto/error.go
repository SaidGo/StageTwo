package dto

// Error — стандартная форма ошибки для HTTP-ответов.
// Совместима с существующим кодом, который ожидает поле "Error".
type Error struct {
	Error string `json:"error"`          // текст ошибки (ожидаемое поле роутером)
	Code  int    `json:"code,omitempty"` // опциональный код
}
