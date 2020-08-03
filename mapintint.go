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

import "math"

// MapIntInt can be built by simply calling
// make(MapIntInt).
type MapIntInt map[int]int

// Increment will add 1 to the integer value found
// at the provided key.
func (m MapIntInt) Increment(key int) {
	v, ok := m[key]
	if !ok {
		m[key] = 1
	} else {
		m[key] = v + 1
	}
}

// Values will return the slice of all values
// found in the map.
func (m MapIntInt) Values() (values []int) {
	for _, v := range m {
		values = append(values, v)
	}

	return
}

// Keys will return the slice of all keys
// found in the map.
func (m MapIntInt) Keys() (keys []int) {
	for k := range m {
		keys = append(keys, k)
	}

	return
}

// MaxValue will return he largest value and it's key.
// The key is the first result returned. And the value
// is the second result returned.
func (m MapIntInt) MaxValue() (key int, value int) {
	max := math.MinInt64
	for k, v := range m {
		if v > max {
			max = v
			key = k
			value = v
		}
	}

	return
}

// MinValue will return the smallest value and it's key.
// The key is the first result returned. And the value
// is the second result returned.
func (m MapIntInt) MinValue() (key int, value int) {
	min := math.MaxInt64
	for k, v := range m {
		if v < min {
			min = v
			key = k
			value = v
		}
	}

	return
}

// Contains will check if the map contains a key.
func (m MapIntInt) Contains(key int) bool {
	_, ok := m[key]
	return ok
}
