//go:build wireinject
// +build wireinject

package app

import (
	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web"
	"example.com/local/Go2part/pkg/postgres"

	"github.com/google/wire"
)

func InitLegalEntityHandler(dsn string) *web.LegalEntityHandler {
	wire.Build(
		postgres.NewSQLiteConnection,
		legalentities.NewGormRepository,
		wire.Bind(new(legalentities.Repository), new(*legalentities.GormRepository)),
		legalentities.NewService,
		web.NewLegalEntityHandler,
	)
	return nil
}
