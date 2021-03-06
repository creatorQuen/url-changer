package httphandlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"url-changer/app"
)

type urlGenerator struct {
	service app.KeyGenerator
}

type UrlToCut struct {
	LongUrl string `json:"long_url"`
}

func NewUrlGenerator(service app.KeyGenerator) *urlGenerator {
	return &urlGenerator{service: service}
}

func (u *urlGenerator) GetUrl(ctx echo.Context) error {
	key := ctx.Param("key")
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "key invalid")
	}

	url, err := u.service.GetURL(key)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.Redirect(http.StatusMovedPermanently, url)
}

func (u *urlGenerator) UrlCutter(ctx echo.Context) error {
	var urlToCut UrlToCut
	err := ctx.Bind(&urlToCut)
	if err != nil {
		log.Error("failed to bind UrlCutter data: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	key, err := u.service.MakeKey(urlToCut.LongUrl)

	if err != nil {
		log.Error("service.MakeKey: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, key)
}
