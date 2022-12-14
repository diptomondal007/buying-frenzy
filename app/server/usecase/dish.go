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
	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/model"
	"github.com/diptomondal007/buying-frenzy/app/server/repository"
)

// DishUseCase is interface for dish use case
type DishUseCase interface {
	SearchDish(term string) ([]*common.MenuResp, error)
}

// dishUseCase ...
type dishUseCase struct {
	repo repository.DishRepository
}

// NewDishUseCase returns a new dish use case instance
func NewDishUseCase(repo repository.DishRepository) DishUseCase {
	return &dishUseCase{repo: repo}
}

// SearchDish ...
func (d dishUseCase) SearchDish(term string) ([]*common.MenuResp, error) {
	rs, err := d.repo.SearchDish(term)
	if err != nil {
		return nil, err
	}
	return toDishList(rs), nil
}

// toDishList ...
func toDishList(ds []*model.Dish) []*common.MenuResp {
	resp := make([]*common.MenuResp, 0)
	for i := range ds {
		resp = append(resp, &common.MenuResp{
			ID:    ds[i].ID,
			Name:  ds[i].Name,
			Price: ds[i].Price,
		})
	}
	return resp
}
