// Copyright 2020 Hummility AI Incorporated, All Rights Reserved.
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

// MapFloat64Int uses float64s as keys
// and integers as values.
type MapFloat64Int map[float64]int

// Max will return the key with the largest
// integer value.
func (m MapFloat64Int) Max() float64 {
	max := math.MinInt64
	var returnValue float64
	for value, count := range m {
		if count > max {
			returnValue = value
			max = count
		}
	}

	return returnValue
}

// Increment will add 1 to the integer value
// of the provided float64 key.
func (m MapFloat64Int) Increment(f float64) {
	v, ok := m[f]
	if !ok {
		m[f] = 1
		return
	}
	v++

	m[f] = v
}

// AverageCount will iterate over the map and
// return the average integer value.
func (m MapFloat64Int) AverageCount() int {
	var avg int

	for _, count := range m {
		avg += count
	}
	avg = avg / len(m)

	return avg
}
