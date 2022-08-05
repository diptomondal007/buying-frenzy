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

	"github.com/diptomondal007/buying-frenzy/app/common"
)

func Test_parseOpeningHours(t *testing.T) {
	type args struct {
		openHour string
	}
	tests := []struct {
		name string
		args args
		want []*schedule
	}{

		{
			name: "t-01",
			args: args{openHour: "Mon - Weds 10:15 am - 7:45 pm / Thurs 5:30 pm - 1 am /"},
			want: []*schedule{
				{
					fromWeekday: (*weekDay)(common.ToIntP(1)),
					toWeekDay:   (*weekDay)(common.ToIntP(3)),
					from: &clock{
						hour:   10,
						minute: 15,
					},
					to: &clock{
						hour:   19,
						minute: 45,
					},
				},
				{
					fromWeekday: (*weekDay)(common.ToIntP(4)),
					toWeekDay:   nil,
					from: &clock{
						hour:   17,
						minute: 30,
					},
					to: &clock{
						hour:   1,
						minute: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOpeningHours(tt.args.openHour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOpeningHours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumber(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t-01",
			args: args{'T'},
			want: false,
		},
		{
			name: "t-02",
			args: args{'1'},
			want: true,
		},
		{
			name: "t-03",
			args: args{'a'},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.r); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSchedule(t *testing.T) {
	type args struct {
		openHour string
	}
	tests := []struct {
		name string
		args args
		want *schedule
	}{
		{
			name: "t-01",
			args: args{openHour: "Mon - Tues 7:45 am - 11:15 am"},
			want: &schedule{
				fromWeekday: (*weekDay)(common.ToIntP(1)),
				toWeekDay:   (*weekDay)(common.ToIntP(2)),
				from: &clock{
					hour:   7,
					minute: 45,
				},
				to: &clock{
					hour:   11,
					minute: 15,
				},
			},
		},
		{
			name: "t-02",
			args: args{openHour: "Tues 1:30 pm - 2:30 pm"},
			want: &schedule{
				fromWeekday: (*weekDay)(common.ToIntP(2)),
				toWeekDay:   nil,
				from: &clock{
					hour:   13,
					minute: 30,
				},
				to: &clock{
					hour:   14,
					minute: 30,
				},
			},
		},
		{
			name: "t-03",
			args: args{openHour: "Thurs, Sun 8 am - 1:15 am"},
			want: &schedule{
				fromWeekday: (*weekDay)(common.ToIntP(4)),
				toWeekDay:   (*weekDay)(common.ToIntP(0)),
				from: &clock{
					hour:   8,
					minute: 0,
				},
				to: &clock{
					hour:   1,
					minute: 15,
				},
			},
		},
		{
			name: "t-04",
			args: args{openHour: "Mon - Weds 10:15 am - 7:45 pm"},
			want: &schedule{
				fromWeekday: (*weekDay)(common.ToIntP(1)),
				toWeekDay:   (*weekDay)(common.ToIntP(3)),
				from: &clock{
					hour:   10,
					minute: 15,
				},
				to: &clock{
					hour:   19,
					minute: 45,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSchedule(tt.args.openHour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strToHourMin(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "t-01",
			args:  args{s: "1:00"},
			want:  1,
			want1: 0,
		},
		{
			name:  "t-02",
			args:  args{s: "1"},
			want:  1,
			want1: 0,
		},
		{
			name:  "t-03",
			args:  args{s: ""},
			want:  0,
			want1: 0,
		},
		{
			name:  "t-04",
			args:  args{s: "asd:10"},
			want:  0,
			want1: 0,
		},
		{
			name:  "t-05",
			args:  args{s: "10:p"},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := strToHourMin(tt.args.s)
			if got != tt.want {
				t.Errorf("strToHourMin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("strToHourMin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_weekDay_String(t *testing.T) {
	tests := []struct {
		name string
		w    weekDay
		want string
	}{
		{
			name: "t-01",
			w:    weekDay(0),
			want: "Sunday",
		},
		{
			name: "t-02",
			w:    weekDay(1),
			want: "Monday",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTime(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "t-01",
			args:  args{t: "1:00"},
			want:  0,
			want1: 0,
		},
		{
			name:  "t-02",
			args:  args{t: "1"},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseTime(tt.args.t)
			if got != tt.want {
				t.Errorf("parseTime() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseTime() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
