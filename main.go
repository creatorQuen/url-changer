package main

import (
	"github.com/labstack/echo/v4"
	"url-changer/app"
	"url-changer/infrastructure/httphandlers"
	"url-changer/infrastructure/localservices"
	repo2 "url-changer/infrastructure/repo"
)

func main() {
	e := echo.New()

	repo := repo2.NewKeySaver()

	geyGeneratorService := localservices.NewKeyGenerator()

	service := app.NewUrlCutterService(*geyGeneratorService, repo)

	handler := httphandlers.NewUrlGenerator(service)

	e.POST("/urlcutter", handler.UrlCutter)
	e.GET("/:key", handler.GetUrl)
	e.Logger.Fatal(e.Start(":8088"))
}
