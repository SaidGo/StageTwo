package app

import (
	"context"
	"github.com/sirupsen/logrus"
)

func (a *App) Seed(ctx context.Context) error {
	logrus.Info(">>> SEED STARTED <<<")

	// Здесь можно подключать нужные сиды, если App включает соответствующие сервисы.
	// Например:
	// err := a.LegalEntitiesService.Seed(ctx)
	// if err != nil {
	//     return err
	// }

	logrus.Info(">>> SEED COMPLETED <<<")
	return nil
}
