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
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound            = errors.New("resource not found")
	ErrInvalidPage         = errors.New("invalid page request")
	ErrNotAcceptable       = errors.New("not acceptable")
	ErrConflict            = errors.New("data conflict or already exist")
	ErrBadRequest          = errors.New("bad request, check param or body")
	ErrInternalServerError = errors.New("internal server response")
)

func getStatusCode(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInvalidPage:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotAcceptable:
		return http.StatusNotAcceptable
	default:
		wrapErr := &WrapErr{}
		if errors.As(err, wrapErr) {
			return wrapErr.StatusCode
		}
		return http.StatusInternalServerError
	}
}

// RespondError takes an `response` and a `customErr message` args
// to log the response to system and return to client
func RespondError(err error, customErr ...error) (int, Response) {
	combinedErr := err
	statusCode := getStatusCode(err)
	resp := Response{Success: false, Message: err.Error(), StatusCode: statusCode}
	if len(customErr) > 0 {
		resp.Message = customErr[0].Error()
		combinedErr = fmt.Errorf("%s : %s", combinedErr, customErr[0])
	}
	if statusCode == http.StatusInternalServerError {
		resp.Message = "something went wrong"
	}
	return statusCode, resp
}

type WrapErr struct {
	StatusCode int
	ErrCode    string
	Err        error
}

// implements response interface
func (e WrapErr) Error() string {
	return e.Err.Error()
}

// Unwrap the errors.Unwrap interface
func (e WrapErr) Unwrap() error {
	return e.Err // Returns inner response
}

// WrapError returns wrapped error object
func WrapError(err error, statusCode int, errCode string) error {
	return WrapErr{
		Err:        err,
		ErrCode:    errCode,
		StatusCode: statusCode,
	}
}
