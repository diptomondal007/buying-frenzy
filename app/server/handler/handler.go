package handler

import (
	"github.com/labstack/echo/v4"
)

type handler struct {
	e *echo.Echo
}

func NewHandler(e *echo.Echo) {
	h := handler{e: e}

	// restaurant group
	rg := e.Group("/api/v1/restaurant")
	rg.GET("/open", h.openRestaurants)
	rg.GET("/list", h.list)
	//rg.GET("/search")
	//
	//ug := e.Group("/api/v1/user")
	//ug.POST("/purchase")
	//
	//dg := e.Group("/api/v1/dishes")
	//dg.GET("/search")
}
