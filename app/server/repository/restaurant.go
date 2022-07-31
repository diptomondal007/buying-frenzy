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
	"github.com/diptomondal007/buying-frenzy/app/common/model"
	"github.com/jmoiron/sqlx"
)

type restaurant struct {
	db *sqlx.DB
}

func (r restaurant) GetOpenRestaurants() ([]*model.Restaurant, error) {
	res := make([]*model.Restaurant, 0)

	if err := r.db.Select(&res, `select * from restaurant`); err != nil {
		return nil, err
	}
	return res, nil
}

type Restaurant interface {
	GetOpenRestaurants() ([]*model.Restaurant, error)
}

func NewRestaurantRepo(db *sqlx.DB) Restaurant {
	return &restaurant{db: db}
}
