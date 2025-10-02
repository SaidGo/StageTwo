//go:build disable_extras

package gates

import "example.com/local/Go2part/internal/app"

type dictionaryRepo interface {
	GetUserFederatons(userUUID string) []string
	FindFederation(uuid string) (interface{}, bool)
	GetUserCompanies(userUUID string) []string
}

type Service struct {
	*app.BaseService

	dict            dictionaryRepo
	commentsLimit   int
	federationLimit int
}
