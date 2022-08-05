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

package response

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRespondError(t *testing.T) {
	type args struct {
		err       error
		customErr []error
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RespondError(tt.args.err, tt.args.customErr...)
			if got != tt.want {
				t.Errorf("RespondError() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RespondError() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWrapErr_Error(t *testing.T) {
	type fields struct {
		StatusCode int
		ErrCode    string
		Err        error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := WrapErr{
				StatusCode: tt.fields.StatusCode,
				ErrCode:    tt.fields.ErrCode,
				Err:        tt.fields.Err,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapErr_Unwrap(t *testing.T) {
	type fields struct {
		StatusCode int
		ErrCode    string
		Err        error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := WrapErr{
				StatusCode: tt.fields.StatusCode,
				ErrCode:    tt.fields.ErrCode,
				Err:        tt.fields.Err,
			}
			if err := e.Unwrap(); (err != nil) != tt.wantErr {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWrapError(t *testing.T) {
	type args struct {
		err        error
		statusCode int
		errCode    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WrapError(tt.args.err, tt.args.statusCode, tt.args.errCode); (err != nil) != tt.wantErr {
				t.Errorf("WrapError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getStatusCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t-01",
			args: args{err: ErrBadRequest},
			want: 400,
		},
		{
			name: "t-02",
			args: args{err: fmt.Errorf("unexpected error")},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatusCode(tt.args.err); got != tt.want {
				t.Errorf("getStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
