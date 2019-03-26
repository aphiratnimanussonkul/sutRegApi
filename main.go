package main

import (
	"log"
	"os"

	"github.com/fooku/sutRegApi/pkg/api"
	"github.com/fooku/sutRegApi/pkg/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	// mongoURL = "mongodb://course:Test1234@ds243212.mlab.com:43212/course"
	mongoURL = "mongodb://127.0.0.1:27017/course"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8100"
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	api.Mount(e)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	e.Logger.Fatal(e.Start(":" + port))
}
