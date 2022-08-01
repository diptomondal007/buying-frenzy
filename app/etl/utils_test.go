package etl

import (
	"reflect"
	"testing"
)

func Test_parseOpeningHours(t *testing.T) {
	type args struct {
		openHour string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "t-01",
			args: args{openHour: "Mon - Weds 10:15 am - 7:45 pm / Thurs 5:30 pm - 1 am / Fri 3:45 pm - 7:30 pm / Sat 7:45 am - 8:15 am / Sun 7:30 am - 8 am"},
			want: []string{"Mon - Weds 10:15 am - 7:45 pm", "Thurs 5:30 pm - 1 am", "Fri 3:45 pm - 7:30 pm", "Sat 7:45 am - 8:15 am", "Sun 7:30 am - 8 am"},
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

func Test_isLetter(t *testing.T) {
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
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
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
				fromWeekday: toIntP(1),
				toWeekDay:   toIntP(2),
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
				fromWeekday: toIntP(2),
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
				fromWeekday: toIntP(4),
				toWeekDay:   toIntP(0),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSchedule(tt.args.openHour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toIntP(a int) *weekDay {
	return (*weekDay)(&a)
}
