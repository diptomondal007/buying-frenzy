package response

import (
	"reflect"
	"testing"
)

func TestRespondSuccess(t *testing.T) {
	type args struct {
		statusCode int
		message    string
		data       interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 Response
	}{
		{
			name: "t-01",
			args: args{
				statusCode: 200,
				message:    "success",
				data:       []interface{}{struct{ name string }{name: "dipto"}},
			},
			want: 200,
			want1: Response{
				Success:    true,
				Message:    "success",
				StatusCode: 200,
				Data:       []interface{}{struct{ name string }{name: "dipto"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RespondSuccess(tt.args.statusCode, tt.args.message, tt.args.data)
			if got != tt.want {
				t.Errorf("RespondSuccess() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RespondSuccess() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
