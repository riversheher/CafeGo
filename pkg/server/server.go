package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"github.com/rainbowriverrr/CafeGo/internal/controller"
)

type Server struct {
	db *sql.DB
}

func registerControllers(e *echo.Echo) {
	controller.RegisterIngredientCTL(e)
	controller.RegisterMenuCTL(e)
	controller.RegisterProductCTL(e)
	controller.RegisterOrderCTL(e)
	controller.RegisterUserCTL(e)
}

func (srv *Server) Start() {
	e := echo.New()

	registerControllers(e)
	e.Logger.Fatal(e.Start(":1323"))
}
