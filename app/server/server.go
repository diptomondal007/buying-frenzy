// Licensed to Dipto Mondal under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Dipto Mondal licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package server

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/diptomondal007/buying-frenzy/app/server/handler"
	"github.com/diptomondal007/buying-frenzy/app/server/repository"
	"github.com/diptomondal007/buying-frenzy/app/server/usecase"
	"github.com/diptomondal007/buying-frenzy/infrastructure/conn"
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

	err := conn.ConnectDB()
	if err != nil {
		log.Println("db connection unsuccessful! error: ", err)
		os.Exit(1)
	}

	rp := repository.NewRestaurantRepo(conn.GetDB().DB)
	us := usecase.NewRestaurantUseCase(rp)
	handler.NewHandler(e, us)

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
