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

// UserEntry is to read user data from stream. We don't want to load the whole file in memory because
// that may cause memory spike and may overflow total memory allocated for the app
type UserEntry struct {
	User  User
	Error error
}
