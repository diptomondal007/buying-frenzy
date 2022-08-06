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

func TestSearchRestaurantSuccessful(t *testing.T) {
	p := response.Response{
		Success:    true,
		Message:    "request successful!",
		StatusCode: 200,
		Data: []common.RestaurantResp{
			{
				ID:   "00301558-f4eb-499f-9e62-fb087bc4bd80",
				Name: "Afternoon Tea at the Briarwood Inn",
			},
		},
	}

	j, _ := json.Marshal(p)

	s := echo.New()

	q := make(url.Values)
	q.Set("q", "Tea")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)

	db, mock := common.MockSqlxDB()
	defer db.Close()

	rr := repository.NewRestaurantRepo(db)

	ru := usecase.NewRestaurantUseCase(rr)

	query := `SELECT "r"."id", "r"."name" FROM "restaurant" AS "r" WHERE (SIMILARITY(name, 'Tea') > 0.2) ORDER BY SIMILARITY(name, 'Tea') DESC`
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("00301558-f4eb-499f-9e62-fb087bc4bd80", "Afternoon Tea at the Briarwood Inn")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	h := NewHandler(s, ru, nil, nil)

	if assert.NoError(t, h.search(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(j)+"\n", rec.Body.String())
	}
}

func TestListRestaurantSuccessful(t *testing.T) {
	p := response.Response{
		Success:    true,
		Message:    "request successful!",
		StatusCode: 200,
		Data: []common.RestaurantResp{
			{
				ID:   "00301558-f4eb-499f-9e62-fb087bc4bd80",
				Name: "Afternoon Tea at the Briarwood Inn",
			},
		},
	}

	j, _ := json.Marshal(p)

	s := echo.New()

	q := make(url.Values)
	q.Set("price_low", "10")
	q.Set("price_high", "200")
	q.Set("more_than", "1")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)

	db, mock := common.MockSqlxDB()
	defer db.Close()

	rr := repository.NewRestaurantRepo(db)

	ru := usecase.NewRestaurantUseCase(rr)

	query := `SELECT "r"."id", "r"."name" FROM "restaurant" AS "r" INNER JOIN "dish" AS "di" ON ("r"."id" = "di"."restaurant_id") WHERE (("price" <= 200) AND ("price" >= 10)) GROUP BY "r"."id" HAVING (COUNT("r"."id") > 1) ORDER BY "r"."name" ASC`
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("00301558-f4eb-499f-9e62-fb087bc4bd80", "Afternoon Tea at the Briarwood Inn")

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	h := NewHandler(s, ru, nil, nil)

	if assert.NoError(t, h.list(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(j)+"\n", rec.Body.String())
	}
}

func TestListRestaurantBadRequest(t *testing.T) {
	p := response.Response{
		Success:    false,
		Message:    "low or high value of price range is missing",
		StatusCode: 400,
	}

	j, _ := json.Marshal(p)

	s := echo.New()

	q := make(url.Values)
	q.Set("price_low", "10")
	q.Set("more_than", "1")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)

	db, _ := common.MockSqlxDB()
	defer db.Close()

	rr := repository.NewRestaurantRepo(db)

	ru := usecase.NewRestaurantUseCase(rr)

	h := NewHandler(s, ru, nil, nil)

	if assert.NoError(t, h.list(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(j)+"\n", rec.Body.String())
	}
}
