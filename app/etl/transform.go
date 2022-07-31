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

import "github.com/diptomondal007/buying-frenzy/app/common"

type Transformer interface {
	transformUserData(user *common.User) error
}

type transformer struct {
	loader Loader
}

func (t transformer) transformUserData(user *common.User) error {
	//TODO implement me
	return nil
}

func newTransformer(loader Loader) Transformer {
	return transformer{loader: loader}
}
