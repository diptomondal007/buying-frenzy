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
