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

package model

import "time"

type PGTime time.Time

const (
	PGTimeFormat = "15:04:05"
	// PGTimeHOURMINFormat = "15:04"
)

const (
	USERInfoTable        = "user_info"
	PURCHASEHistoryTable = "purchase_history"
	RESTAURANTTable      = "restaurant"
	DISHTable            = "dish"
	OPENHourTable        = "open_hour"
)

// UserInfo is the model for db user_info
type UserInfo struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	CashBalance float64 `db:"cash_balance"`
}

// Restaurant is the model for restaurant
type Restaurant struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	CashBalance float64 `db:"cash_balance"`
}

type OpenHour struct {
	ID           int     `db:"-"`
	WeekName     string  `db:"week_name"`
	StartTime    *PGTime `db:"start_time"`
	ClosingTime  *PGTime `db:"closing_time"`
	RestaurantID string  `db:"restaurant_id"`
}

type Dish struct {
	ID           string  `db:"id"`
	Name         string  `db:"name"`
	Price        float64 `db:"price"`
	RestaurantID string  `db:"restaurant_id"`
}

type PurchaseHistory struct {
	ID                int       `db:"-"`
	TransactionAmount float64   `db:"transaction_amount"`
	TransactionDate   time.Time `db:"transaction_date"`
	RestaurantID      string    `db:"restaurant_id"`
	DishID            string    `db:"dish_id"`
	UserID            int       `db:"user_id"`
}

type RestaurantWithDish struct {
	RestaurantID string `db:"r_id"`
	DishID       string `db:"d_id"`
	//DishName     string  `db:"d_name"`
	DishPrice float64 `db:"d_price"`
}
