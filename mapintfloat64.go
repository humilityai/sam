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

// MapIntFloat64 uses integers as keys
// and float64s as values.
type MapIntFloat64 map[int]float64

// Min will return the integer key with the smallest
// float64 value.
func (m MapIntFloat64) Min() (index int) {
	min := math.MaxFloat64
	for i, v := range m {
		if v < min {
			min = v
			index = i
		}
	}

	return
}

// Max will return the integer key with the
// largest float64 value.
func (m MapIntFloat64) Max() (index int) {
	max := float64(math.MinInt64)
	for i, v := range m {
		if v < max {
			max = v
			index = i
		}
	}

	return
}

// MaxN will return a MapIntFloat64 containing the N largest
// key/value pairs by value.
func (m MapIntFloat64) MaxN(n int) MapIntFloat64 {
	maxN := make(MapIntFloat64)

	for i, v := range m {
		if len(maxN) < n {
			maxN[i] = v
		} else {
			minIndex := maxN.Min()
			if v > maxN[minIndex] {
				delete(maxN, minIndex)
				maxN[i] = v
			}
		}
	}

	return maxN
}
