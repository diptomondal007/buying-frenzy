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

package etl

import (
	"log"
	"path/filepath"

	"github.com/diptomondal007/buying-frenzy/app/common"
	"github.com/diptomondal007/buying-frenzy/pkg/json_steam"
)

// Args contains necessary info for extracting data
type Args struct {
	Directory      string
	UserData       string
	RestaurantData string
}

// ETL is type for etl
type ETL struct {
	Args        Args
	Transformer Transformer
}

// NewETL a new instance of ETL
func NewETL(args Args) *ETL {
	return &ETL{
		Args:        args,
		Transformer: newTransformer(newLoader()),
	}
}

// Run is the main run function of etl sub command
func (e *ETL) Run() error {
	return e.extract()
}

func (e *ETL) extract() error {
	s := json_steam.NewJSONStreamer()
	go s.Start(filepath.Join(e.Args.Directory, e.Args.UserData))

	go func() {
		for _ = range s.Want() {
			s.Value() <- &common.User{}
		}
	}()

	for data := range s.Watch() {
		if data.Error != nil {
			log.Println(data.Error)
			return data.Error
		}

		u, ok := data.Data.(*common.User)
		if !ok {
			log.Println("not ok")
		}

		err := e.Transformer.transformUserData(u)
		if err != nil {
			log.Println("error while transforming data >>>> ", err)
			return nil
		}
	}
	return nil
}
