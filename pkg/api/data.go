package api

import (
	"net/http"

	"github.com/fooku/sutRegApi/pkg/model"
	"github.com/fooku/sutRegApi/pkg/sprape"
	"github.com/labstack/echo"
)

func getCurse(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	courseid, _ := model.GetCourseid(id)

	if courseid == "" {
		return c.NoContent(http.StatusOK)
	}
	_ = sprape.GetDataReg(courseid, c)

	return nil
	//fmt.Println(data)

	// var content struct {
	// 	Courseid string `json:"courseid"`
	// }
	// content.Courseid = courseid

}
