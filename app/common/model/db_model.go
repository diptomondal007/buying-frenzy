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

package common

import "time"

type PGTime time.Time

const (
	PGTimeFormat        = "15:04:05"
	PGTimeHOURMINFormat = "15:04"
)

// UserInfo is the model for db user_info
type UserInfo struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	CashBalance float64 `db:"cash_balance"`
}

// Restaurant is the model for restaurant
type Restaurant struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	CashBalance float64 `db:"cash_balance"`
}

type OpenHour struct {
	id           int
	WeekName     string
	StartTime    PGTime
	ClosingTime  PGTime
	RestaurantID int `db:"restaurant_id"`
}
