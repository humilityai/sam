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

// SliceBool is a slice/array of boolean values.
type SliceBool []bool

// Equal will check if SliceBool is equal to provided
// object.
func (s SliceBool) Equal(element Slice) bool {
	input, ok := element.(SliceBool)
	if !ok {
		return false
	}

	if len(input) != len(s) {
		return false
	}

	for i, v := range s {
		if v != input[i] {
			return false
		}
	}

	return true
}

// Len will return the length
// of the boolean slice as an integer.
func (s SliceBool) Len() int {
	return len(s)
}

// Get is just a function version of what can be done
// more easily with `[index]`.
func (s SliceBool) Get(index int) interface{} {
	return s[index]
}

// Set is just a function version of what can be done
// more easily with `[index]`.
func (s SliceBool) Set(index int, input interface{}) {
	v, ok := input.(bool)
	if ok {
		s[index] = v
	}
}

// Subslice is just to satisfy the slice interface.
func (s SliceBool) Subslice(start, end int) Slice {
	return SliceBool(s[start:end])
}

// Type will return the type name of the values
// stored
func (s SliceBool) Type() string {
	return BoolType
}

// FalseIndices will return the list of all indices
// in the slice that have a false value.
func (s SliceBool) FalseIndices() SliceInt {
	var indices SliceInt
	for i, v := range s {
		if !v {
			indices = append(indices, i)
		}
	}

	return indices
}

// FalseCount will return the total number of
// false values in the boolean slice.
func (s SliceBool) FalseCount() (count int) {
	for _, v := range s {
		if !v {
			count++
		}
	}

	return
}

// FalsePercentage will return the percentage of values
// that are false.
func (s SliceBool) FalsePercentage() float64 {
	return float64(s.FalseCount()) / float64(len(s))
}

// TrueIndices will return the list of all indices
// in the slice that have a true value.
func (s SliceBool) TrueIndices() SliceInt {
	var indices SliceInt
	for i, v := range s {
		if v {
			indices = append(indices, i)
		}
	}

	return indices
}

// TrueCount will return the total number of
// true values in the boolean slice.
func (s SliceBool) TrueCount() (count int) {
	for _, v := range s {
		if v {
			count++
		}
	}

	return
}

// TruePercentage will return the percentage of values
// that are true.
func (s SliceBool) TruePercentage() float64 {
	return float64(s.TrueCount()) / float64(len(s))
}
