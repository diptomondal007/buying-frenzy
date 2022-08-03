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
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	"github.com/labstack/echo/v4"

	"github.com/diptomondal007/buying-frenzy/app/common"
)

//
func (h *handler) openRestaurants(c echo.Context) error {
	now := time.Now()

	if c.QueryParam("date_time") != "" {
		n, err := dateparse.ParseAny(c.QueryParam("date_time"))
		if err != nil {
			h.e.Logger.Error("error while parsing input date time", err)
			return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "bad format of date time"})
		}
		now = n
		log.Println("now", now)
	}

	rs, err := h.rc.ListRestaurantsByFilter(now)
	if err != nil {
		log.Println("error while fetching from db. error: ", err)
		return c.JSON(http.StatusInternalServerError, common.ErrResp{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, common.Resp{
		Message: "request successful!",
		Data:    rs,
	})
}

func (h *handler) list(c echo.Context) error {
	lowP := c.QueryParam("price_low")
	highP := c.QueryParam("price_high")

	if lowP == "" || highP == "" {
		return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "low or high value of price range is missing!"})
	}

	lessThan := c.QueryParam("less_than")
	moreThan := c.QueryParam("more_than")
	if lessThan == "" && moreThan == "" {
		return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "both more_than and less_than param can't be empty!"})
	}

	if lessThan != "" && moreThan != "" {
		return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "both more_than and less_than param exist! please use one of them at a time"})
	}

	lowPF, err := strconv.ParseFloat(lowP, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "invalid value for low range"})
	}

	highPF, err := strconv.ParseFloat(highP, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "invalid value for high range"})
	}

	var lessThanV *int
	var moreThanV *int

	if lessThan != "" {
		v, err := strconv.Atoi(lessThan)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "invalid value for 'less_than' param"})
		}
		lessThanV = &v
	}

	if moreThan != "" {
		v, err := strconv.Atoi(moreThan)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "invalid value for 'more than' param"})
		}
		moreThanV = &v
	}

	rs, err := h.rc.ListRestaurantsByDishFilter(common.RestaurantFilter{
		PriceRange: common.PriceRange{
			High: highPF,
			Low:  lowPF,
		},
		QuantityRange: common.QuantityRange{
			MoreThan: moreThanV,
			LessThan: lessThanV,
		},
	})
	if err != nil {
		log.Println("error while fetching from db. error: ", err)
		return c.JSON(http.StatusInternalServerError, common.ErrResp{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, common.Resp{
		Message: "request successful!",
		Data:    rs,
	})
}

func (h *handler) search(c echo.Context) error {
	q := c.QueryParam("q")

	if q == "" {
		return c.JSON(http.StatusBadRequest, common.ErrResp{Error: "search term missing!"})
	}

	rs, err := h.rc.SearchRestaurant(q)
	if err != nil {
		log.Println("error while fetching from db. error: ", err)
		return c.JSON(http.StatusInternalServerError, common.ErrResp{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, common.Resp{
		Message: "request successful!",
		Data:    rs,
	})
}
