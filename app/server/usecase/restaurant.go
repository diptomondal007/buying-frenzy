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
	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/model"
	"github.com/diptomondal007/buying-frenzy/app/server/repository"
)

type RestaurantUseCase interface {
	ListRestaurantsByFilter() ([]*common.Restaurant, error)
}

type restaurantUseCase struct {
	repo repository.Restaurant
}

func (r restaurantUseCase) ListRestaurantsByFilter() ([]*common.Restaurant, error) {
	rs, err := r.repo.GetOpenRestaurants()
	if err != nil {
		return nil, err
	}
	return toRestaurantList(rs), nil
}

func NewRestaurantUseCase(repo repository.Restaurant) RestaurantUseCase {
	return &restaurantUseCase{repo: repo}
}

func toRestaurantList(rs []*model.Restaurant) []*common.Restaurant {
	resp := make([]*common.Restaurant, 0)
	for i := range rs {
		resp = append(resp, &common.Restaurant{
			ID:   rs[i].ID,
			Name: rs[i].Name,
		})
	}
	return resp
}
