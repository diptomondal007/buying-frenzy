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
	"reflect"
	"testing"

	"github.com/diptomondal007/buying-frenzy/app/common/model"
)

func Test_toOpenHours(t *testing.T) {
	type args struct {
		rid string
		h   string
	}
	tests := []struct {
		name string
		args args
		want []model.OpenHour
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{
				rid: "d12b245e-0d11-4e74-9586-25e5460283b6",
				h:   "Mon - Tues 7:45 am - 11:15 am",
			},
			want: []model.OpenHour{
				{
					WeekName:     "Monday",
					StartTime:    model.NewPGTimeFromHourMinute(7, 45),
					ClosingTime:  model.NewPGTimeFromHourMinute(7, 45),
					RestaurantID: "d12b245e-0d11-4e74-9586-25e5460283b6",
				},
				{
					WeekName:     "Tuesday",
					StartTime:    model.NewPGTimeFromHourMinute(7, 45),
					ClosingTime:  model.NewPGTimeFromHourMinute(7, 45),
					RestaurantID: "d12b245e-0d11-4e74-9586-25e5460283b6",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toOpenHours(tt.args.rid, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toOpenHours() = %v, want %v", got, tt.want)
			}
		})
	}
}
