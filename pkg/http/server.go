package http

import (
	"github.com/labstack/echo/v4"
	"github.com/torfstack/shrt/pkg/service"
)

func StartServer(s service.ShrtService) {
	e := echo.New()
	e.GET("/:something", Shrt(s))
	e.Logger.Fatal(e.Start(":8080"))
}
