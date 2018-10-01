package main

import (
	"log"
	"os"

	"github.com/fooku/sutRegApi/pkg/api"
	"github.com/fooku/sutRegApi/pkg/model"
	"github.com/labstack/echo"
)

const (
	mongoURL = "mongodb://course:Test1234@ds243212.mlab.com:43212/course"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	e := echo.New()
	api.Mount(e)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	e.Logger.Fatal(e.Start(":" + port))
}
