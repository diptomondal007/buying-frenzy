package server

import (
	"context"
	"github.com/diptomondal007/buying-frenzy/app/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
)

func attach(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
}

type Server struct {
	server *echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	handler.NewHandler(e)
	// attaching middleware to echo server
	attach(e)
	return &Server{server: e}
}

func (s *Server) Run() {
	go func() {
		s.server.Logger.Error(s.server.Start(":8000"))
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	log.Println("-> shutting down server gracefully ....")

	err := s.server.Shutdown(context.Background())
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("√ successfully shut down!")
}
