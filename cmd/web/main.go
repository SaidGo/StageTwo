package main

import (
	"example.com/local/Go2part/internal/app"
	"example.com/local/Go2part/internal/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	a := app.NewApp("legalentities.db")

	router.RegisterLegalEntityRoutes(e, a.LegalEntityHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
