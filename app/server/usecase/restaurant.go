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

package usecase

import (
	"time"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/model"
	"github.com/diptomondal007/buying-frenzy/app/server/repository"
)

type RestaurantUseCase interface {
	ListRestaurantsByFilter(datetime time.Time) ([]*common.RestaurantResp, error)
	ListRestaurantsByDishFilter(filter common.RestaurantFilter) ([]*common.RestaurantResp, error)
}

type restaurantUseCase struct {
	repo repository.Restaurant
}

func (r restaurantUseCase) ListRestaurantsByFilter(datetime time.Time) ([]*common.RestaurantResp, error) {
	weekDay := datetime.Weekday().String()

	rs, err := r.repo.GetOpenRestaurants(datetime.Hour(), datetime.Minute(), weekDay)
	if err != nil {
		return nil, err
	}
	return toRestaurantList(rs), nil
}

func (r restaurantUseCase) ListRestaurantsByDishFilter(filter common.RestaurantFilter) ([]*common.RestaurantResp, error) {
	rs, err := r.repo.GetRestaurantsByDishFilter(filter)
	if err != nil {
		return nil, err
	}
	return toRestaurantList(rs), nil
}

func NewRestaurantUseCase(repo repository.Restaurant) RestaurantUseCase {
	return &restaurantUseCase{repo: repo}
}

func toRestaurantList(rs []*model.Restaurant) []*common.RestaurantResp {
	resp := make([]*common.RestaurantResp, 0)
	for i := range rs {
		resp = append(resp, &common.RestaurantResp{
			ID:   rs[i].ID,
			Name: rs[i].Name,
		})
	}
	return resp
}
