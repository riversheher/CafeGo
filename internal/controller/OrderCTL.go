package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterOrderCTL(e *echo.Echo) {
	e.GET("/orders", GetOrders)
	e.GET("/orders/:id", GetOrder)
	e.POST("/orders", CreateOrder)
	e.PUT("/orders/:id", UpdateOrder)
	e.DELETE("/orders/:id", DeleteOrder)
	e.GET("/cart", GetCart)
}

func GetOrders(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetOrders")
}

func GetOrder(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetOrder")
}

func CreateOrder(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateOrder")
}

func UpdateOrder(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateOrder")
}

func DeleteOrder(c echo.Context) error {
	return c.JSON(http.StatusOK, "DeleteOrder")
}

func GetCart(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetCart")
}
