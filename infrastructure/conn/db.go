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

package conn

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/diptomondal007/buying-frenzy/infrastructure/config"
)

type postgresClient struct {
	*sqlx.DB
}

var (
	db *postgresClient
)

func GetDB() *postgresClient {
	return db
}

func ConnectDB() error {
	if db != nil {
		log.Println("db already initialized!")
		return nil
	}
	cfg := config.Get().DB

	d, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name))
	if err != nil {
		return err
	}

	db = &postgresClient{d}
	return nil
}
