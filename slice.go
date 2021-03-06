// Copyright 2020 Humility AI Incorporated, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sam

import (
	"fmt"
	"reflect"
)

var (
	BoolType    = fmt.Sprint(reflect.TypeOf(true))
	Float64Type = fmt.Sprint(reflect.TypeOf(float64(1)))
	StringType  = fmt.Sprint(reflect.TypeOf(""))
	Int64Type   = fmt.Sprint(reflect.TypeOf(int64(1)))
)

// Slice is a generic interface that
// can be used to represent any slice
// type.
// Current primary use is for external packages.
type Slice interface {
	Equal(interface{}) bool
	Get(int) interface{}
	Len() int
	Set(int, interface{})
	Type() string
	Subslice(start, end int) Slice
}
