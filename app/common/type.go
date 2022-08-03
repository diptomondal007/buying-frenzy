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

// User holds the structure for user
type User struct {
	ID              int               `json:"id"`
	Name            string            `json:"name"`
	CashBalance     float64           `json:"cashBalance"`
	PurchaseHistory []PurchaseHistory `json:"purchaseHistory"`
}

// PurchaseHistory is the single purchase history of a user
type PurchaseHistory struct {
	DishName          string  `json:"dishName"`
	RestaurantName    string  `json:"restaurantName"`
	TransactionAmount float64 `json:"transactionAmount"`
	TransactionDate   string  `json:"transactionDate"`
}

type RestaurantResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Restaurant struct {
	RestaurantName string  `json:"restaurantName"`
	CashBalance    float64 `json:"cashBalance"`
	OpeningHours   string  `json:"openingHours"`
	Menu           []Menu  `json:"menu"`
}

type Menu struct {
	DishName string  `json:"dishName"`
	Price    float64 `json:"price"`
}

type RestaurantFilter struct {
	PriceRange    PriceRange
	QuantityRange QuantityRange
}

type PriceRange struct {
	High float64
	Low  float64
}

type MenuResp struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type QuantityRange struct {
	MoreThan *int
	LessThan *int
}

type UserResp struct {
	Balance float64 `json:"current_balance"`
}

type PurchaseRequest struct {
	RestaurantID string `json:"restaurant_id"`
	MenuID       string `json:"menu_id"`
}
