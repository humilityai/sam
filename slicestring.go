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

import (
	"sort"
	"strings"
	"unicode"
)

// SliceString is a slice/array of strings.
type SliceString []string

// Equal will check if the string slice is equal to
// the supplied object argument and return a boolean.
// It can check for equality against any (arbitrary) argument.
func (s1 SliceString) Equal(input interface{}) bool {
	s2, ok := input.(SliceString)
	if !ok {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	for _, value := range s1 {
		if !s2.Contains(value) {
			return false
		}
	}

	return true
}

// Contains will check if the slice contains the supplied
// string argument.
func (s1 SliceString) Contains(s string) bool {
	for _, v := range s1 {
		if v == s {
			return true
		}
	}

	return false
}

// ToLower will lowercase all the strings in the slice.
func (s1 SliceString) ToLower() {
	for i, v := range s1 {
		s1[i] = strings.ToLower(v)
	}
}

// String will convert the string slice into a single string,
// with each string being delimited by the supplied delimiter.
func (s1 SliceString) String(delimeter string) string {
	return strings.Join(s1, delimeter)
}

// SortedString will lexicographically sort the strings
// in the slice and return them as a single string delimited
// by the supplied delimieter argument.
func (s1 SliceString) SortedString(delimeter string) string {
	sort.Sort(s1)
	return strings.Join(s1, delimeter)
}

// Len ...
func (s1 SliceString) Len() int {
	return len(s1)
}

// Swap ...
func (s1 SliceString) Swap(i, j int) {
	s1[i], s1[j] = s1[j], s1[i]
}

// Less provides lexicographic sorting.
func (s1 SliceString) Less(i, j int) bool {
	iRunes := []rune(s1[i])
	jRunes := []rune(s1[j])

	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		// the lowercase runes are the same, so compare the original
		if ir != jr {
			return ir < jr
		}
	}

	return false
}
