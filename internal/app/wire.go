//go:build wireinject
// +build wireinject

package app

import (
	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web/olegalentity"

	"github.com/google/wire"
)

var appSet = wire.NewSet(
	NewDB,
	NewApp,
	NewRouter,

	legalentities.NewRepository,
	legalentities.NewService,
	wire.Bind(new(legalentities.ServiceInterface), new(*legalentities.Service)),

	olegalentity.NewLegalEntityHandler,
	wire.Bind(new(olegalentity.ServerInterface), new(*olegalentity.LegalEntityHandler)),
)

func InitApp() (*App, error) {
	wire.Build(appSet)
	return nil, nil
}
