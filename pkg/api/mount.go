package api

import (
	"github.com/labstack/echo"
)

// Mount xx
func Mount(e *echo.Echo) {
	e.GET("/getcourse/:id", getCurse)
	e.POST("/data", getData)
	e.GET("/data", outData)
	e.GET("/data2", outData2)
	e.GET("/data3", outData3)
	e.GET("/data4", outData4)
	e.GET("/data5", outData5)
	e.GET("/data6", outData6)
}
