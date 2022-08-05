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

package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/common/response"
	"github.com/diptomondal007/buying-frenzy/app/server/repository"
	"github.com/diptomondal007/buying-frenzy/app/server/usecase"
)

func TestSearchDishSuccessful(t *testing.T) {
	p := response.Response{
		Success:    true,
		Message:    "request successful!",
		StatusCode: 200,
		Data: []common.MenuResp{
			{
				ID:    "0d9625a0-1298-40f2-b168-167b4ad70d74",
				Name:  "Pie",
				Price: 100.1,
			},
		},
	}

	j, _ := json.Marshal(p)

	s := echo.New()

	q := make(url.Values)
	q.Set("q", "piz")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)

	db, mock := common.MockSqlxDB()
	defer db.Close()

	dr := repository.NewDishRepo(db)

	du := usecase.NewDishUseCase(dr)

	query := `SELECT "d"."id", "d"."name", "d"."price" FROM "dish" AS "d" WHERE (SIMILARITY(name, 'piz') > 0.2) ORDER BY SIMILARITY(name, 'piz') DESC`
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow("0d9625a0-1298-40f2-b168-167b4ad70d74", "Pie", 100.10)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	h := NewHandler(s, nil, du, nil)

	if assert.NoError(t, h.searchDish(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(j)+"\n", rec.Body.String())
	}
}

func TestSearchDishBadRequest(t *testing.T) {
	p := response.Response{
		Success:    false,
		Message:    "query param 'q' required",
		StatusCode: 400,
	}

	j, _ := json.Marshal(p)

	s := echo.New()

	//q := make(url.Values)
	//q.Set("q", "piz")
	req := httptest.NewRequest(http.MethodGet, "/?", nil)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)

	db, _ := common.MockSqlxDB()
	defer db.Close()

	dr := repository.NewDishRepo(db)

	du := usecase.NewDishUseCase(dr)

	h := NewHandler(s, nil, du, nil)

	if assert.NoError(t, h.searchDish(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(j)+"\n", rec.Body.String())
	}
}
