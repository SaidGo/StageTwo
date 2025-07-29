//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web/olegalentity"
)

var appSet = wire.NewSet(
	NewApp,
	NewRouter,
	NewDB,
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