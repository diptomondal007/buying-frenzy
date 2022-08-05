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
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/diptomondal007/buying-frenzy/app/common/model"
	"github.com/diptomondal007/buying-frenzy/app/common/response"
)

// userRepository ...
type userRepository struct {
	db *sqlx.DB
}

// UserRepository ...
type UserRepository interface {
	PurchaseDish(userID int, restaurantID string, menuID string) (model.UserInfo, error)
}

// NewUserRepo returns a new user repo instance
func NewUserRepo(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) PurchaseDish(userID int, restaurantID, menuID string) (model.UserInfo, error) {
	var updatedUser model.UserInfo

	tx := u.db.MustBegin()
	defer tx.Rollback()

	var user model.UserInfo
	q, _, err := goqu.From(goqu.T(model.USERInfoTable).As("u")).
		Select("u.*").
		Where(
			goqu.Ex{"id": goqu.Op{"eq": userID}},
		).ToSQL()
	if err != nil {
		return model.UserInfo{}, err
	}

	if err := tx.Get(&user, q); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserInfo{}, response.WrapError(
				fmt.Errorf("user not found"),
				http.StatusNotFound, "")
		}
		return model.UserInfo{}, err
	}

	var rDish model.RestaurantWithDish
	q, _, err = goqu.From(goqu.T(model.RESTAURANTTable).As("r")).
		Select(goqu.I("r.id").As("r_id"), goqu.I("d.id").As("d_id"), goqu.I("d.price").As("d_price")).
		Join(goqu.T(model.DISHTable).As("d"),
			goqu.On(goqu.I("r.id").Eq(goqu.I("d.restaurant_id")))).
		Where(
			goqu.Ex{"r.id": goqu.Op{"eq": restaurantID}},
			goqu.Ex{"d.id": goqu.Op{"eq": menuID}},
		).ToSQL()

	if err != nil {
		return model.UserInfo{}, err
	}

	if err := tx.Get(&rDish, q); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserInfo{}, response.WrapError(
				fmt.Errorf("restaurant or dish does not exist"),
				http.StatusNotFound, "")
		}
		return model.UserInfo{}, err
	}
	if rDish.DishPrice > user.CashBalance {
		return model.UserInfo{}, response.WrapError(
			fmt.Errorf("you don't have enough cash to buy this dish! you have $%.2f", user.CashBalance),
			http.StatusNotAcceptable, "")
	}

	q, _, err = goqu.Update(model.RESTAURANTTable).Set(map[string]interface{}{
		"cash_balance": goqu.L("cash_balance + ?", rDish.DishPrice),
	}).Where(goqu.Ex{"id": goqu.Op{"eq": restaurantID}}).ToSQL()
	if err != nil {
		return model.UserInfo{}, err
	}
	tx.MustExec(q)

	q, _, err = goqu.Update(model.USERInfoTable).Set(map[string]interface{}{
		"cash_balance": goqu.L("cash_balance - ?", rDish.DishPrice),
	}).Where(goqu.Ex{"id": goqu.Op{"eq": userID}}).ToSQL()
	if err != nil {
		return model.UserInfo{}, err
	}
	tx.MustExec(q)

	q, _, err = goqu.Insert(goqu.T(model.PURCHASEHistoryTable)).Rows(model.PurchaseHistory{
		TransactionAmount: rDish.DishPrice,
		TransactionDate:   time.Now().UTC(),
		RestaurantID:      rDish.RestaurantID,
		DishID:            rDish.DishID,
		UserID:            userID,
	}).ToSQL()
	if err != nil {
		return model.UserInfo{}, err
	}

	_, err = tx.Exec(q)
	if err != nil {
		log.Println(err)
		return model.UserInfo{}, err
	}

	q, _, err = goqu.From(goqu.T(model.USERInfoTable).As("u")).
		Select("u.*").Where(
		goqu.Ex{"id": goqu.Op{"eq": userID}},
	).ToSQL()
	if err != nil {
		return model.UserInfo{}, err
	}

	if err := tx.Get(&updatedUser, q); err != nil {
		return model.UserInfo{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.UserInfo{}, err
	}

	return updatedUser, nil
}
