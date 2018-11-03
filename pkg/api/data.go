package api

import (
	"net/http"

	"github.com/fooku/sutRegApi/pkg/model"
	"github.com/fooku/sutRegApi/pkg/sprape"
	"github.com/labstack/echo"
)

func getCurse(c echo.Context) error {
	chRes1 := make(chan error)

	// User ID from path `users/:id`
	id := c.QueryParam("id")
	semester := c.QueryParam("semester")
	acadyear := c.QueryParam("acadyear")
	//id := c.Param("id")

	courseid, _ := model.GetCid(id)

	if courseid == "" {
		return c.NoContent(http.StatusOK)
	}

	go func() {
		chRes1 <- sprape.GetDataReg(courseid, c, semester, acadyear)
	}()

	return <-chRes1
	//fmt.Println(data)

	// var content struct {
	// 	Courseid string `json:"courseid"`
	// }
	// content.Courseid = courseid

}
