package olegalentity

import "example.com/local/Go2part/internal/legalentities"

// NewLegalEntityHandler возвращает хендлер с "внедрённым" сервисом.
func NewLegalEntityHandler(svc *legalentities.Service) *LegalEntityHandler {
	return &LegalEntityHandler{svc: svc}
}
