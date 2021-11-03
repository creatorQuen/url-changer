package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"url-changer/app"
	"url-changer/infrastructure/httphandlers"
	"url-changer/infrastructure/localservices"
	"url-changer/infrastructure/repo"
)

func main() {
	e := echo.New()

	repo := repo.NewKeySaver()

	keyGeneratorService := localservices.NewKeyGenerator()
	service := app.NewUrlCutterService(*keyGeneratorService, repo)

	handler := httphandlers.NewUrlGenerator(service)

	e.POST("/urlcutter", handler.UrlCutter)
	e.GET("/:key", handler.GetUrl)
	e.Logger.Fatal(e.Start(":8088"))
}
