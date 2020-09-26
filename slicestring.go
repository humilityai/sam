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
	"sort"
	"strings"
	"unicode"
)

// SliceString is a slice/array of strings.
type SliceString []string

// Equal will check if the string slice is equal to
// the supplied object argument and return a boolean.
// It can check for equality against any (arbitrary) argument.
func (s SliceString) Equal(input interface{}) bool {
	s2, ok := input.(SliceString)
	if !ok {
		return false
	}

	if len(s) != len(s2) {
		return false
	}

	for _, value := range s {
		if !s2.Contains(value) {
			return false
		}
	}

	return true
}

func (s SliceString) Type() string {
	return StringType
}

func (s SliceString) Get(idx int) interface{} {
	return s[idx]
}

func (s SliceString) Set(idx int, value interface{}) {
	v, ok := value.(string)
	if !ok {
		return
	}
	s[idx] = v
}

func (s SliceString) Subslice(start, end int) Slice {
	return s[start:end]
}

// Contains will check if the slice contains the supplied
// string argument.
func (s SliceString) Contains(input string) bool {
	for _, v := range s {
		if v == input {
			return true
		}
	}

	return false
}

// ToLower will lowercase all the strings in the slice.
func (s SliceString) ToLower() {
	for i, v := range s {
		s[i] = strings.ToLower(v)
	}
}

// String will convert the string slice into a single string,
// with each string being delimited by the supplied delimiter.
func (s SliceString) String(delimeter string) string {
	return strings.Join(s, delimeter)
}

// SortedString will lexicographically sort the strings
// in the slice and return them as a single string delimited
// by the supplied delimieter argument.
func (s SliceString) SortedString(delimeter string) string {
	sort.Sort(s)
	return strings.Join(s, delimeter)
}

// Len ...
func (s SliceString) Len() int {
	return len(s)
}

// Swap ...
func (s SliceString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less provides lexicographic sorting.
// Original code: https://stackoverflow.com/a/35087122
func (s SliceString) Less(i, j int) bool {
	iRunes := []rune(s[i])
	jRunes := []rune(s[j])

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
