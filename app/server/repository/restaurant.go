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

package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/model"
)

type restaurant struct {
	db *sqlx.DB
}

type Restaurant interface {
	GetOpenRestaurants(hour, minute int, weekDay string) ([]*model.Restaurant, error)
	GetRestaurantsByDishFilter(f common.RestaurantFilter) ([]*model.Restaurant, error)
	SearchRestaurant(term string) ([]*model.Restaurant, error)
}

func NewRestaurantRepo(db *sqlx.DB) Restaurant {
	return &restaurant{db: db}
}

func (r restaurant) GetOpenRestaurants(hour, minute int, weekDay string) ([]*model.Restaurant, error) {
	res := make([]*model.Restaurant, 0)
	t := model.NewPGTimeFromHourMinute(hour, minute)

	sql, _, err := goqu.From(goqu.T(model.RESTAURANTTable).As("r")).
		Select("r.id", "r.name").
		Join(goqu.T(model.OPENHourTable).As("oh"), goqu.On(goqu.Ex{"r.id": goqu.I("oh.restaurant_id")})).
		Where(goqu.Ex{
			"oh.week_name": goqu.Op{"eq": weekDay},
		}).GroupBy("r.id", "oh.closing_time", "oh.start_time", "oh.week_name").
		Having(goqu.Or(
			goqu.Ex{
				"oh.start_time":   goqu.Op{"lte": t},
				"oh.closing_time": goqu.Op{"gte": t},
			},
			goqu.Or(
				goqu.And(
					goqu.Ex{
						"oh.closing_time": goqu.Op{"lt": goqu.I("oh.start_time")},
					},
					goqu.Ex{
						"oh.start_time":   goqu.Op{"lte": t},
						"oh.closing_time": goqu.Op{"lte": t},
					}),
			),
		)).ToSQL()

	if err != nil {
		return nil, err
	}

	if err := r.db.Select(&res, sql); err != nil {
		return nil, err
	}
	return res, nil
}

func (r restaurant) GetRestaurantsByDishFilter(f common.RestaurantFilter) ([]*model.Restaurant, error) {
	res := make([]*model.Restaurant, 0)

	d := goqu.From(goqu.T(model.RESTAURANTTable).As("r")).
		Select("r.id", "r.name").
		Join(goqu.T(model.DISHTable).As("di"), goqu.On(goqu.Ex{"r.id": goqu.I("di.restaurant_id")})).
		Where(goqu.And(
			goqu.Ex{
				"price": goqu.Op{"lte": f.PriceRange.High},
			},
			goqu.Ex{
				"price": goqu.Op{"gte": f.PriceRange.Low},
			}),
		).GroupBy("r.id").Order(goqu.I("r.name").Asc())

	if f.QuantityRange.MoreThan != nil {
		d = d.Having(goqu.COUNT(goqu.I("r.id")).Gt(*f.QuantityRange.MoreThan))
	}

	if f.QuantityRange.LessThan != nil {
		d = d.Having(goqu.COUNT(goqu.I("r.id")).Lt(*f.QuantityRange.LessThan))
	}

	sql, _, err := d.ToSQL()
	if err != nil {
		return nil, err
	}

	if err := r.db.Select(&res, sql); err != nil {
		return nil, err
	}
	return res, nil
}

func (r restaurant) SearchRestaurant(term string) ([]*model.Restaurant, error) {
	res := make([]*model.Restaurant, 0)

	d := goqu.From(goqu.T(model.RESTAURANTTable).As("r")).
		Select("r.id", "r.name").
		Where(
			goqu.L("SIMILARITY(name, ?)", term).Gt(0.2),
		).Order(goqu.L("SIMILARITY(name, ?)", term).Desc())

	sql, _, err := d.ToSQL()
	if err != nil {
		return nil, err
	}

	if err := r.db.Select(&res, sql); err != nil {
		return nil, err
	}
	return res, nil
}
