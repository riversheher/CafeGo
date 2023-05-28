package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterProductCTL(e *echo.Echo) {
	e.GET("/products", GetProducts)
	e.GET("/products/:id", GetProduct)
	e.POST("/products", CreateProduct)
	e.PUT("/products/:id", UpdateProduct)
	e.DELETE("/products/:id", DeleteProduct)
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetProducts")
}

func GetProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetProduct")
}

func CreateProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateProduct")
}

func UpdateProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateProduct")
}

func DeleteProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "DeleteProduct")
}
