package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterIngredientCTL(e *echo.Echo) {
	e.GET("/ingredients", GetIngredients)
	e.GET("/ingredients/:id", GetIngredient)
	e.POST("/ingredients", CreateIngredient)
	e.PUT("/ingredients/:id", UpdateIngredient)
	e.DELETE("/ingredients/:id", DeleteIngredient)
}

func GetIngredients(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetIngredients")
}

func GetIngredient(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetIngredient")
}

func CreateIngredient(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateIngredient")
}

func UpdateIngredient(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateIngredient")
}

func DeleteIngredient(c echo.Context) error {
	return c.JSON(http.StatusOK, "DeleteIngredient")
}
