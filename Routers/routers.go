package Routers

import (
	"todo/Handlers"

	"github.com/labstack/echo"
)

func Start() {

	e := echo.New()
	e.GET("/details", Handlers.GetDetails)
	e.GET("/details/:id", Handlers.GetIdDetails)
	e.POST("/details", Handlers.AddDetails)
	e.PUT("/details/:title", Handlers.UpdateDetails)
	e.DELETE("/details/:title", Handlers.DeleteDetails)

	e.Start(":3030")
}
