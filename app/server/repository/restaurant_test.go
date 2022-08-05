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
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/model"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var r = []*model.Restaurant{
	{
		ID:   "0d9625a0-1298-40f2-b168-167b4ad70d74",
		Name: "Pie",
	},
}

func TestSearchRestaurant(t *testing.T) {
	db, mock := common.MockSqlxDB()
	defer db.Close()

	dr := NewRestaurantRepo(db)

	query := `SELECT "r"."id", "r"."name" FROM "restaurant" AS "r" WHERE (SIMILARITY(name, 'piz') > 0.2) ORDER BY SIMILARITY(name, 'piz') DESC`

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("0d9625a0-1298-40f2-b168-167b4ad70d74", "Pie")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	dishes, err := dr.SearchRestaurant("piz")
	assert.NoError(t, err)
	assert.Equal(t, r, dishes)
}

func TestGetOpenRestaurant(t *testing.T) {
	db, mock := common.MockSqlxDB()
	defer db.Close()

	dr := NewRestaurantRepo(db)

	query := `SELECT "r"."id", "r"."name" FROM "restaurant" AS "r" INNER JOIN "open_hour" AS "oh" ON ("r"."id" = "oh"."restaurant_id") WHERE ("oh"."week_name" = 'Sunday') GROUP BY "r"."id", "oh"."closing_time", "oh"."start_time", "oh"."week_name" HAVING ((("oh"."closing_time" >= '10:45:00') AND ("oh"."start_time" <= '10:45:00')) OR (("oh"."closing_time" < "oh"."start_time") AND (("oh"."closing_time" <= '10:45:00') AND ("oh"."start_time" <= '10:45:00'))))`

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("0d9625a0-1298-40f2-b168-167b4ad70d74", "Pie")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	dishes, err := dr.GetOpenRestaurants(10, 45, "Sunday")
	assert.NoError(t, err)
	assert.Equal(t, r, dishes)
}

func TestGetRestaurantsByDishFilter(t *testing.T) {
	db, mock := common.MockSqlxDB()
	defer db.Close()

	dr := NewRestaurantRepo(db)

	query := `SELECT "r"."id", "r"."name" FROM "restaurant" AS "r" INNER JOIN "dish" AS "di" ON ("r"."id" = "di"."restaurant_id") WHERE (("price" <= 200) AND ("price" >= 100)) GROUP BY "r"."id" HAVING (COUNT("r"."id") > 1) ORDER BY "r"."name" ASC`

	rows := sqlmock.NewRows([]string{"id", "name"})

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	dishes, err := dr.GetRestaurantsByDishFilter(common.RestaurantFilter{
		PriceRange: common.PriceRange{
			High: 200,
			Low:  100,
		},
		QuantityRange: common.QuantityRange{
			MoreThan: common.ToIntP(1),
			LessThan: nil,
		},
	})
	assert.NoError(t, err)
	assert.Len(t, dishes, 0)
}
