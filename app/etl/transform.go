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

package etl

import (
	"time"

	"github.com/araddon/dateparse"
	uuid "github.com/satori/go.uuid"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/model"
)

type Transformer interface {
	transformUserData(user *common.User) error
	transformRestaurantData(r *common.Restaurant) error
}

type transformer struct {
	loader Loader
	dishM  map[string]string
	restM  map[string]string
}

func (t transformer) transformUserData(user *common.User) error {
	um := model.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		CashBalance: user.CashBalance,
	}

	return t.loader.loadUserData(um, t.toPurchaseHistoryModel(user.ID, user.PurchaseHistory))
}

func (t transformer) transformRestaurantData(restaurant *common.Restaurant) error {
	rid := uuid.NewV4().String()
	r := model.Restaurant{
		ID:          rid,
		Name:        restaurant.RestaurantName,
		CashBalance: restaurant.CashBalance,
	}

	t.restM[restaurant.RestaurantName] = rid

	openHours := toOpenHours(rid, restaurant.OpeningHours)
	mList := t.toMenuList(rid, restaurant.Menu)

	return t.loader.loadRestaurantData(r, openHours, mList)
}

func NewTransformer(loader Loader) Transformer {
	return transformer{
		loader: loader,
		dishM:  make(map[string]string),
		restM:  make(map[string]string),
	}
}

func (t transformer) toPurchaseHistoryModel(userID int, history []common.PurchaseHistory) []model.PurchaseHistory {
	hs := make([]model.PurchaseHistory, 0)

	for i := range history {
		ti, _ := dateparse.ParseIn(history[i].TransactionDate, time.UTC)
		restaurantID := t.restM[history[i].RestaurantName]
		dishID := t.dishM[history[i].DishName]

		hs = append(hs, model.PurchaseHistory{
			TransactionAmount: history[i].TransactionAmount,
			TransactionDate:   ti.UTC(),
			RestaurantID:      restaurantID,
			DishID:            dishID,
			UserID:            userID,
		})
	}
	return hs
}

func (t transformer) toMenuList(rid string, m []common.Menu) []model.Dish {
	d := make([]model.Dish, 0)
	for i := range m {
		id := uuid.NewV4().String()
		t.dishM[m[i].DishName] = id
		d = append(d, model.Dish{
			ID:           id,
			Name:         m[i].DishName,
			Price:        m[i].Price,
			RestaurantID: rid,
		})
	}
	return d
}

func toOpenHours(rid, h string) []model.OpenHour {
	openHours := make([]model.OpenHour, 0)
	schedules := parseOpeningHours(h)

	for i := range schedules {
		s := schedules[i]
		from := 0
		to := 0
		if s.fromWeekday != nil {
			from = int(*s.fromWeekday)
		}

		if s.toWeekDay == nil {
			openHours = append(openHours, model.OpenHour{
				WeekName:     s.fromWeekday.String(),
				StartTime:    model.NewPGTimeFromHourMinute(s.from.hour, s.from.minute),
				ClosingTime:  model.NewPGTimeFromHourMinute(s.to.hour, s.to.minute),
				RestaurantID: rid,
			})
			continue
		}

		to = int(*s.toWeekDay)
		temp := 0

		// swap
		if from > to {
			temp = to
			to = from
			from = temp
		}

		for f := from; f <= to; f++ {
			openHours = append(openHours, model.OpenHour{
				WeekName:     longDayNames[f],
				StartTime:    model.NewPGTimeFromHourMinute(s.from.hour, s.from.minute),
				ClosingTime:  model.NewPGTimeFromHourMinute(s.to.hour, s.to.minute),
				RestaurantID: rid,
			})
		}
	}
	return openHours
}
