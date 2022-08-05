package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/diptomondal007/buying-frenzy/app/server/repository"
	"github.com/diptomondal007/buying-frenzy/app/server/usecase"
)

func TestPurchasedSuccessful(t *testing.T) {
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

	//query := `SELECT * from user_info where id =0`
	//rows := sqlmock.NewRows([]string{"id", "name", "cash_balance"}).
	//AddRow(0, "Test", 100.10)

	mock.ExpectBegin()
	//mock.ExpectQuery(e).ExpectQuery(query).WillReturnRows(rows)

	h := NewHandler(s, nil, nil, us)

	if assert.NoError(t, h.purchase(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		//assert.Equal(t, purchaseBody, rec.Body.String())
	}
}
