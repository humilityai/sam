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

// SliceInt64 is a slice/array of integers
type SliceInt64 []int64

// Equal will only return true if the object is a SliceInt64
// of the same length and has smae values in the same positions.
func (s SliceInt64) Equal(input interface{}) bool {
	sl, ok := input.(SliceInt64)
	if !ok {
		return false
	}

	if len(sl) != len(s) {
		return false
	}

	for i, v := range s {
		if sl[i] != v {
			return false
		}
	}

	return true
}

// Get will return the int64 value at the specified index
func (s SliceInt64) Get(idx int) interface{} {
	return s[idx]
}

// Len will return the length of the SliceInt64
func (s SliceInt64) Len() int {
	return len(s)
}

// Set will set the provided value at the provided index
func (s SliceInt64) Set(idx int, value interface{}) {
	i, ok := value.(int64)
	if ok {
		s[idx] = i
	}
}

// Type will return an Int64Type value for SliceInt64 objects
func (s SliceInt64) Type() string {
	return Int64Type
}

// Subslice will return a new SliceInt64 that is a subslice of
// this SliceInt64
func (s SliceInt64) Subslice(start, end int) Slice {
	return s[start:end]
}

// Product will return the total product of
// all values in the slice.
func (s SliceInt64) Product() int64 {
	var product int64
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
func (s SliceInt64) Sum() int64 {
	var sum int64

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
func (s SliceInt64) Min() (index int, value int64) {
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
func (s SliceInt64) Max() (index int) {
	max := int64(math.MinInt64)
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
func (s SliceInt64) Contains(i int64) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

// IncrementPosition will add 1 to the integer value
// found at the supplied index argument.
func (s SliceInt64) IncrementPosition(index int) {
	if index > len(s)-1 || index < 0 {
		return
	}

	s[index]++
}
