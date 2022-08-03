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
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/diptomondal007/buying-frenzy/app/common"
)

func (h *handler) searchDish(c echo.Context) error {
	q := c.QueryParam("q")

	if q == "" {
		return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "search term missing!"})
	}

	rs, err := h.dc.SearchDish(q)
	if err != nil {
		log.Println("error while fetching from db. error: ", err)
		return c.JSON(http.StatusInternalServerError, common.ErrResp{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, common.Resp{
		Message: "request successful!",
		Data:    rs,
	})
}
