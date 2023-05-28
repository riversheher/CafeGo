package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMenuCTL(e *echo.Echo) {
	e.GET("/menus", GetMenus)
	e.GET("/menus/:id", GetMenu)
	e.POST("/menus", CreateMenu)
	e.PUT("/menus/:id", UpdateMenu)
	e.DELETE("/menus/:id", DeleteMenu)
}

func GetMenus(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetMenus")
}

func GetMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetMenu")
}

func CreateMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateMenu")
}

func UpdateMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateMenu")
}

func DeleteMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, "DeleteMenu")
}
