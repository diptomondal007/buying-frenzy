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

	"github.com/diptomondal007/buying-frenzy/app/common/model"
)

// dishRepository ...
type dishRepository struct {
	db *sqlx.DB
}

// DishRepository ...
type DishRepository interface {
	SearchDish(term string) ([]*model.Dish, error)
}

// NewDishRepo returns a new dish repo instance
func NewDishRepo(db *sqlx.DB) DishRepository {
	return &dishRepository{db: db}
}

// SearchDish is the repo for searching dish in db
func (dr dishRepository) SearchDish(term string) ([]*model.Dish, error) {
	dishes := make([]*model.Dish, 0)

	d := goqu.From(goqu.T(model.DISHTable).As("d")).
		Select("d.id", "d.name", "d.price").
		Where(
			goqu.L("SIMILARITY(name, ?)", term).Gt(0.2),
		).Order(goqu.L("SIMILARITY(name, ?)", term).Desc())

	sql, _, err := d.ToSQL()
	if err != nil {
		return nil, err
	}

	if err := dr.db.Select(&dishes, sql); err != nil {
		return nil, err
	}
	return dishes, nil
}
