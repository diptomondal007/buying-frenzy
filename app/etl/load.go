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
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/diptomondal007/buying-frenzy/app/common/model"
)

type Loader interface {
	loadUserData(user model.UserInfo, histories []model.PurchaseHistory) error
	loadRestaurantData(r model.Restaurant, s []model.OpenHour, m []model.Dish) error
}

type loader struct {
	db *sqlx.DB
}

func (l loader) loadUserData(user model.UserInfo, histories []model.PurchaseHistory) error {
	tx := l.db.MustBegin()

	sql, _, err := goqu.Insert(model.USERInfoTable).Rows(user).ToSQL()
	if err != nil {
		return err
	}
	tx.MustExec(sql)
	log.Println(sql)

	if len(histories) > 0 {
		sql, _, err = goqu.Insert(model.PURCHASEHistoryTable).Rows(histories).ToSQL()
		if err != nil {
			return err
		}
		tx.MustExec(sql)
		log.Println(sql)
	}
	return tx.Commit()
}

func (l loader) loadRestaurantData(r model.Restaurant, s []model.OpenHour, m []model.Dish) error {
	tx := l.db.MustBegin()
	sql, _, err := goqu.Insert(model.RESTAURANTTable).Rows(&r).ToSQL()
	if err != nil {
		return err
	}
	tx.MustExec(sql)

	sql, _, err = goqu.Insert(model.OPENHourTable).Rows(s).ToSQL()
	log.Println(sql)
	if err != nil {
		return err
	}
	tx.MustExec(sql)

	sql, _, err = goqu.Insert(model.DISHTable).Rows(m).ToSQL()
	if err != nil {
		return err
	}
	tx.MustExec(sql)
	return tx.Commit()
}

func NewLoader(db *sqlx.DB) Loader {
	return &loader{
		db: db,
	}
}
