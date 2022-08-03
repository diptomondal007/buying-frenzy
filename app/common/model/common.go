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

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func NewPGTimeFromHourMinute(hour, minute int) *PGTime {
	pg := PGTime(time.Date(0, 0, 0, hour, minute, 0, 0, time.UTC))
	return &pg
}

func (t *PGTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = PGTime(v)
	case nil:
		*t = PGTime{}
	default:
		return fmt.Errorf("cannot sql.Scan() MyTime from: %#v", v)
	}
	return nil
}

func (t PGTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(PGTimeFormat)), nil
}

func (t *PGTime) UnmarshalText(value string) error {
	dd, err := time.Parse(PGTimeFormat, value)
	if err != nil {
		return err
	}
	*t = PGTime(dd)
	return nil
}

func (PGTime) GormDataType() string {
	return "TIME"
}

func (t PGTime) String() string {
	return time.Time(t).Format(PGTimeFormat)
}

func (t *PGTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, errors.New("got nil")
	}
	return json.Marshal(t.String())
}

func (t *PGTime) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	return t.UnmarshalText(s)
}

func (t *PGTime) Time() *time.Time {
	tt := time.Time(*t)
	return &tt
}
