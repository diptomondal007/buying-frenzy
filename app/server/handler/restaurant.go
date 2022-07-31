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
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/diptomondal007/buying-frenzy/app/common"
)

func (h *handler) openRestaurants(c echo.Context) error {
	now := time.Now()

	if c.QueryParam("date_time") != "" {
		n, err := time.Parse(time.RFC3339, c.QueryParam("date_time"))
		if err != nil {
			h.e.Logger.Error("error while parsing input date time", err)
			return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "bad format of date time"})
		}
		now = n
	}

	return c.JSON(http.StatusOK, common.Resp{
		Message: "request successful!",
		Data:    now,
	})
}

func (h *handler) list(c echo.Context) error {

	return c.JSON(http.StatusOK, common.Resp{
		Message: "request successful!",
		Data:    "",
	})
}
