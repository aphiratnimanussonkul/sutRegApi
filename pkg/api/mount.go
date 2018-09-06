package api

import (
	"github.com/labstack/echo"
)

// Mount xx
func Mount(e *echo.Echo) {
	e.GET("/getcourse/:id", getCurse)
}
