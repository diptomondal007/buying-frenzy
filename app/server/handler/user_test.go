package handler

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/app/server/repository"
	"github.com/diptomondal007/buying-frenzy/app/server/usecase"
)

func TestPurchasedUnSuccessful(t *testing.T) {
	s := echo.New()

	purchaseBody := `{
    					"restaurant_id": "00017a27-5fcc-4e01-acab-b791aa0a6292",
    					"menu_id": "0d9625a0-1298-40f2-b168-167b4ad70d74"
					}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(purchaseBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)
	c.SetPath("/api/v1/user/purchase/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("0")

	db, mock, err := sqlmock.New()
	if err != nil {
		return
	}

	dbp := sqlx.NewDb(db, "postgres")
	ur := repository.NewUserRepo(dbp)

	us := usecase.NewUserUseCase(ur)

	mock.ExpectBegin()

	h := NewHandler(s, nil, nil, us)

	if assert.NoError(t, h.purchase(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		//assert.Equal(t, purchaseBody, rec.Body.String())
	}
}

func TestPurchasedSuccessful(t *testing.T) {
	s := echo.New()

	res := `{"success":true,"message":"purchased successfully!","status_code":202,"data":{"current_balance":100.1}}`
	purchaseBody := `{
    					"restaurant_id": "00017a27-5fcc-4e01-acab-b791aa0a6292",
    					"menu_id": "0d9625a0-1298-40f2-b168-167b4ad70d74"
					}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(purchaseBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)
	c.SetPath("/api/v1/user/purchase/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("0")

	db, mock := common.MockSqlxDB()
	ur := repository.NewUserRepo(db)

	us := usecase.NewUserUseCase(ur)

	uRows := sqlmock.NewRows([]string{"id", "name", "cash_balance"}).AddRow(0, "Test", 100.10)

	mock.ExpectBegin()
	query := `SELECT "u".* FROM "user_info" AS "u" WHERE ("id" = 0)`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(uRows)

	dRest := mock.NewRows([]string{"r_id", "d_id", "d_price"}).AddRow("00017a27-5fcc-4e01-acab-b791aa0a6292", "0d9625a0-1298-40f2-b168-167b4ad70d74", 12.50)
	query = `SELECT "r"."id" AS "r_id", "d"."id" AS "d_id", "d"."price" AS "d_price" FROM "restaurant" AS "r" INNER JOIN "dish" AS "d" ON ("r"."id" = "d"."restaurant_id") WHERE (("r"."id" = '00017a27-5fcc-4e01-acab-b791aa0a6292') AND ("d"."id" = '0d9625a0-1298-40f2-b168-167b4ad70d74'))`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(dRest)

	query = `UPDATE "restaurant" SET "cash_balance"=cash_balance + 12.5 WHERE ("id" = '00017a27-5fcc-4e01-acab-b791aa0a6292')`

	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(0, 1))

	query = `UPDATE "user_info" SET "cash_balance"=cash_balance - 12.5 WHERE ("id" = 0)`
	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(0, 1))

	query = `INSERT INTO "purchase_history" ("dish_id", "restaurant_id", "transaction_amount", "transaction_date", "user_id")`
	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewErrorResult(nil))

	query = `SELECT "u".* FROM "user_info" AS "u" WHERE ("id" = 0)`
	uRows = sqlmock.NewRows([]string{"id", "name", "cash_balance"}).AddRow(0, "Test", 100.10)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(uRows)

	mock.ExpectCommit()

	//mock.ExpectQuery(e).ExpectQuery(query).WillReturnRows(rows)

	h := NewHandler(s, nil, nil, us)

	if assert.NoError(t, h.purchase(c)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
		assert.Equal(t, res+"\n", rec.Body.String())
	}
}

func TestPurchasedUnSuccessfulForBalance(t *testing.T) {
	s := echo.New()

	res := `{"success":false,"message":"you don't have enough cash to buy this dish! you have $10.10","status_code":406}`
	purchaseBody := `{
    					"restaurant_id": "00017a27-5fcc-4e01-acab-b791aa0a6292",
    					"menu_id": "0d9625a0-1298-40f2-b168-167b4ad70d74"
					}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(purchaseBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := s.NewContext(req, rec)
	c.SetPath("/api/v1/user/purchase/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues("0")

	db, mock := common.MockSqlxDB()
	ur := repository.NewUserRepo(db)

	us := usecase.NewUserUseCase(ur)

	uRows := sqlmock.NewRows([]string{"id", "name", "cash_balance"}).AddRow(0, "Test", 10.10)

	mock.ExpectBegin()
	query := `SELECT "u".* FROM "user_info" AS "u" WHERE ("id" = 0)`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(uRows)

	dRest := mock.NewRows([]string{"r_id", "d_id", "d_price"}).AddRow("00017a27-5fcc-4e01-acab-b791aa0a6292", "0d9625a0-1298-40f2-b168-167b4ad70d74", 12.50)
	query = `SELECT "r"."id" AS "r_id", "d"."id" AS "d_id", "d"."price" AS "d_price" FROM "restaurant" AS "r" INNER JOIN "dish" AS "d" ON ("r"."id" = "d"."restaurant_id") WHERE (("r"."id" = '00017a27-5fcc-4e01-acab-b791aa0a6292') AND ("d"."id" = '0d9625a0-1298-40f2-b168-167b4ad70d74'))`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(dRest)

	mock.ExpectRollback()

	h := NewHandler(s, nil, nil, us)

	if assert.NoError(t, h.purchase(c)) {
		assert.Equal(t, http.StatusNotAcceptable, rec.Code)
		assert.Equal(t, res+"\n", rec.Body.String())
	}
}
