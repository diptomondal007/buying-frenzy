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

package usecase

import (
	"github.com/diptomondal007/buying-frenzy/app/common/model"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/server/repository"
)

// UserUseCase ...
type userUseCase struct {
	repo repository.UserRepository
}

// UserUseCase is interface for user use case
type UserUseCase interface {
	PurchaseDish(userID int, restaurantID, menuID string) (*common.UserResp, error)
}

// NewUserUseCase returns a new user use case instance
func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (u userUseCase) PurchaseDish(userID int, restaurantID, menuID string) (*common.UserResp, error) {
	us, err := u.repo.PurchaseDish(userID, restaurantID, menuID)
	if err != nil {
		return nil, err
	}

	return toUserResp(us), nil
}

func toUserResp(info model.UserInfo) *common.UserResp {
	return &common.UserResp{Balance: info.CashBalance}
}
