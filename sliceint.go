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

// SliceInt is a slice/array of integers
type SliceInt []int

// Product will return the total product of
// all values in the slice.
func (s SliceInt) Product() int {
	var product int
	for i, v := range s {
		if i == 0 {
			product = v
			continue
		}
		product *= v
	}

	return product
}

// Sum will return the total sum of
// all values in the slice.
func (s SliceInt) Sum() int {
	var sum int

	for i, v := range s {
		if i == 0 {
			sum = v
			continue
		}
		sum += v
	}

	return sum
}

// Min will return the index and the value of
// the smallest value in the slice.
func (s SliceInt) Min() (index, value int) {
	value = math.MaxInt64
	for i, v := range s {
		if v < value {
			index = i
			value = v
		}
	}

	return
}

// Max will return the index of the largest
// value in the slice.
func (s SliceInt) Max() (index int) {
	max := math.MinInt64
	for i, v := range s {
		if v > max {
			v = max
			index = i
		}
	}

	return
}

// Contains checks if an integer value
// already exists in the slice.
func (s SliceInt) Contains(i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

// IncrementPosition will add 1 to the integer value
// found at the supplied index argument.
func (s SliceInt) IncrementPosition(index int) {
	if index > len(s)-1 || index < 0 {
		return
	}

	s[index]++
}
