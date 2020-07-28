package slices

import (
	"sort"
	"strings"
	"unicode"
)

// SliceString ...
type SliceString []string

// Equal ...
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

// Contains ...
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

// String ...
func (s1 SliceString) String(delimeter string) string {
	return strings.Join(s1, delimeter)
}

// SortedString ...
func (s1 SliceString) SortedString() string {
	sort.Sort(s1)
	return strings.Join(s1, "")
}

// SortedStringWithDelimeter ...
func (s1 SliceString) SortedStringWithDelimeter(delimeter string) string {
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
