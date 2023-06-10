package server

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/rainbowriverrr/CafeGo/internal/controller"

	"sync"
)

type Server struct {
	WaitGroup *sync.WaitGroup
	Ctx       context.Context
}

func registerControllers(e *echo.Echo) {
	controller.RegisterIngredientCTL(e)
	controller.RegisterMenuCTL(e)
	controller.RegisterProductCTL(e)
	controller.RegisterOrderCTL(e)
	controller.RegisterUserCTL(e)
}

func (srv *Server) Start() {
	defer srv.WaitGroup.Done()

	e := echo.New()

	registerControllers(e)
	e.Logger.Fatal(e.Start(":1323"))
}
