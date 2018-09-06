package api

import (
	"fmt"

	"github.com/fooku/sutRegApi/pkg/model"
	"github.com/fooku/sutRegApi/pkg/sprape"
	"github.com/labstack/echo"
)

func getCurse(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	courseid, err := model.GetCourseid(id)
	if err != err {
		return err
	}
	fmt.Println(courseid)

	err = sprape.GetDataReg(courseid, c)

	return err

	//fmt.Println(data)

	// var content struct {
	// 	Courseid string `json:"courseid"`
	// }
	// content.Courseid = courseid

}
