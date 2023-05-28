package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserCTL(e *echo.Echo) {
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUser)
	e.POST("/users", CreateUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetUsers")
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetUser")
}

func CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateUser")
}

func UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateUser")
}

func DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "DeleteUser")
}
