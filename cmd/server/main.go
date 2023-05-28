package main

import (
	"github.com/labstack/echo/v4"

	"github.com/rainbowriverrr/CafeGo/internal/controller"
)

func main() {
	e := echo.New()

	controller.RegisterIngredientCTL(e)

	e.Logger.Fatal(e.Start(":1323"))
}
