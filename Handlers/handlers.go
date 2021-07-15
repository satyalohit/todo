package Handlers

import (
	"fmt"
	"log"
	"net/http"
	"todo/model"
	"todo/server"

	"github.com/labstack/echo"
)

func GetDetails(c echo.Context) error {

	db := server.Database()
	defer db.Close()
	var details []model.Task

	db.Find(&details)

	fmt.Println(details)

	return c.JSON(http.StatusOK, details)

}
func GetIdDetails(c echo.Context) error {

	id := c.Param("id")
	db := server.Database()
	defer db.Close()

	var details model.Task

	db.Find(&details, id)
	if details.Title == "" {
		return c.String(http.StatusOK, "Not There")
	} else {
		return c.JSON(http.StatusOK, details)
	}
}
func AddDetails(c echo.Context) error {
	var details model.Task
	err := c.Bind(&details)
	if err != nil {
		log.Print(err)
	}

	if details.Title == "" {
		return c.String(http.StatusOK, "TITLE should not be null")
	}

	db := server.Database()
	defer db.Close()
	var existingdetails model.Task
	db.Find(&existingdetails, "title=?", details.Title)
	if existingdetails.Title == details.Title {
		return c.String(http.StatusOK, "Details are there")
	}
	db.Create(&details)
	return c.JSON(http.StatusOK, details)
}
func UpdateDetails(c echo.Context) error {

	var details model.Task

	db := server.Database()
	title := c.Param("title")
	defer db.Close()

	db.Find(&details, "title=?", title)
	err := c.Bind(&details)
	if err != nil {
		log.Print(err)
	}
	if details.Title == "" {
		return c.String(http.StatusOK, "TITLE should not be null")
	}
	db.Save(&details)
	return c.JSON(http.StatusOK, details)

}
func DeleteDetails(c echo.Context) error {

	var details model.Task

	db := server.Database()
	title := c.Param("title")
	defer db.Close()

	db.Find(&details, "title=?", title)
	db.Delete(&details)

	return c.String(http.StatusOK, "Successfully deleted")

}
