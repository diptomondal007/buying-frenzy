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

package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/diptomondal007/buying-frenzy/app/server/usecase"
)

type handler struct {
	e  *echo.Echo
	rc usecase.RestaurantUseCase
}

func NewHandler(e *echo.Echo, rc usecase.RestaurantUseCase) {
	h := handler{e: e, rc: rc}

	// restaurant group
	rg := e.Group("/api/v1/restaurant")
	rg.GET("/open", h.openRestaurants)
	rg.GET("/list", h.list)
	rg.GET("/search", h.search)
	//
	//ug := e.Group("/api/v1/user")
	//ug.POST("/purchase")
	//
	//dg := e.Group("/api/v1/dishes")
	//dg.GET("/search")
}
