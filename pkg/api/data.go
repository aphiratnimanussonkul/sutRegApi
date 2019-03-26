package api

import (
	"fmt"
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

func getData(c echo.Context) (err error) {
	u := new(model.Data)
	if err = c.Bind(u); err != nil {
		return
	}
	fmt.Println(u)
	err = model.Insert(u)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func outData(c echo.Context) (err error) {

	err, d := model.Get()

	fmt.Println(err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, d)
}

func outData2(c echo.Context) (err error) {

	err, planet := model.GetCC2()

	fmt.Println(err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, planet)
}

func outData3(c echo.Context) (err error) {

	err, planet := model.GetCC()

	fmt.Println(err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, planet)
}

func outData4(c echo.Context) (err error) {

	err, planet := model.GetCourseData()

	fmt.Println(err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, planet)
}

func outData5(c echo.Context) (err error) {

	err, planet := model.GetCC2()
	err, planet2 := model.GetA()
	fmt.Println(err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]int{
		"count": planet, "count2": planet2,
	})
}

func outData6(c echo.Context) (err error) {

	err, planet := model.GetA()

	fmt.Println(err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]int{
		"count": planet,
	})
}
