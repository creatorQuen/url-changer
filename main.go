package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"url-changer/app"
	"url-changer/infrastructure/httphandlers"
	"url-changer/infrastructure/localservices"
	"url-changer/infrastructure/repository"
)

func main() {
	e := echo.New()
	db := ConnectDB()
	defer db.Close()

	repo := repository.NewKeySaver(db)

	keyGeneratorService := localservices.NewKeyGenerator()
	service := app.NewUrlCutterService(*keyGeneratorService, repo)

	handler := httphandlers.NewUrlGenerator(service)

	e.POST("/urlcutter", handler.UrlCutter)
	e.GET("/:key", handler.GetUrl)
	e.Logger.Fatal(e.Start(":8088"))
}

func ConnectDB() (db *sql.DB) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	return
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "klklklk"
	dbname   = "URLchanger"
)
