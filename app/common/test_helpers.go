package common

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func MockSqlxDB() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil
	}

	dbp := sqlx.NewDb(db, "postgres")
	return dbp, mock
}

func ToIntP(a int) *int {
	return &a
}
