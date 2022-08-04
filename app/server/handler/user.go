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
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/response"
)

func (h *handler) purchase(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		log.Println("bad data given as user id", err)
		return c.JSON(response.RespondError(response.ErrBadRequest, fmt.Errorf("not a valid user id")))
	}

	var pr *common.PurchaseRequest
	err = c.Bind(&pr)
	if err != nil {
		log.Println("bad data given", err)
		return c.JSON(response.RespondError(fmt.Errorf("not a valid request body")))
	}

	if pr.RestaurantID == "" || pr.MenuID == "" {
		log.Println("bad data given", err)
		return c.JSON(response.RespondError(fmt.Errorf("not a valid request body")))
	}

	u, err := h.uc.PurchaseDish(userID, pr.RestaurantID, pr.MenuID)
	if err != nil {
		return c.JSON(response.RespondError(err))
	}

	return c.JSON(response.RespondSuccess(http.StatusAccepted, "purchased successfully!", u))
}
