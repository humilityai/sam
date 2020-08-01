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

import "sort"

// Sortable allows for multiple slices to be ordered
// either ascending or descending according to the order of the
// independent slice.
type Sortable interface {
	Dependents() interface{}
	Independent() interface{}
	Order()
}

// SortStringsByFloat64 ...
type SortStringsByFloat64 struct {
	Dependent SliceString
	Ind       SliceFloat64
	Ascending bool
}

// NewSortStringsByFloat64 will return an object which can have both the string and float64
// slice ordered ascending according to the float64 slice. If ascending is specified as true then
// order will be by float64 ascending, else order will be by float64 descending.
func NewSortStringsByFloat64(dependent SliceString, independent SliceFloat64, ascending bool) *SortStringsByFloat64 {
	return &SortStringsByFloat64{
		Dependent: dependent,
		Ind:       independent,
		Ascending: ascending,
	}
}

// Len ...
func (t *SortStringsByFloat64) Len() int {
	return len(t.Dependent)
}

// Swap ...
func (t *SortStringsByFloat64) Swap(i, j int) {
	t.Dependent[i], t.Dependent[j] = t.Dependent[j], t.Dependent[i]
	t.Ind[i], t.Ind[j] = t.Ind[j], t.Ind[i]
}

// Less ...
func (t *SortStringsByFloat64) Less(i, j int) bool {
	if t.Ascending {
		return t.Ind[i] < t.Ind[j]
	}

	return t.Ind[j] < t.Ind[i]
}

// Dependents ...
func (t *SortStringsByFloat64) Dependents() interface{} {
	return t.Dependent
}

// Independent ...
func (t *SortStringsByFloat64) Independent() interface{} {
	return t.Ind
}

// Order ...
func (t *SortStringsByFloat64) Order() {
	sort.Sort(t)
}

// SortStringsByInt ...
type SortStringsByInt struct {
	Dependent SliceString
	Ind       []int
	Ascending bool
}

// NewSortStringsByInt will return an object which can have both the string and float64
// slice ordered ascending according to the float64 slice. If ascending is specified as true then
// order will be by float64 ascending, else order will be by float64 descending.
func NewSortStringsByInt(dependent SliceString, independent []int, ascending bool) *SortStringsByInt {
	return &SortStringsByInt{
		Dependent: dependent,
		Ind:       independent,
		Ascending: ascending,
	}
}

// Len ...
func (t *SortStringsByInt) Len() int {
	return len(t.Dependent)
}

// Swap ...
func (t *SortStringsByInt) Swap(i, j int) {
	t.Dependent[i], t.Dependent[j] = t.Dependent[j], t.Dependent[i]
	t.Ind[i], t.Ind[j] = t.Ind[j], t.Ind[i]
}

// Less ...
func (t *SortStringsByInt) Less(i, j int) bool {
	if t.Ascending {
		return t.Ind[i] < t.Ind[j]
	}

	return t.Ind[j] < t.Ind[i]
}

// Dependents ...
func (t *SortStringsByInt) Dependents() interface{} {
	return t.Dependent
}

// Independent ...
func (t *SortStringsByInt) Independent() interface{} {
	return t.Ind
}

// Order ...
func (t *SortStringsByInt) Order() {
	sort.Sort(t)
}

// SortSliceFloat64BySliceString ...
type SortSliceFloat64BySliceString struct {
	Dependent SliceFloat64
	Ind       SliceString
	Ascending bool
}

// NewSortSliceFloat64BySliceString ...
func NewSortSliceFloat64BySliceString(dependent SliceFloat64, independent SliceString, ascending bool) *SortSliceFloat64BySliceString {
	return &SortSliceFloat64BySliceString{
		Dependent: dependent,
		Ind:       independent,
		Ascending: ascending,
	}
}

// Len ...
func (t *SortSliceFloat64BySliceString) Len() int {
	return len(t.Dependent)
}

// Swap ...
func (t *SortSliceFloat64BySliceString) Swap(i, j int) {
	t.Dependent[i], t.Dependent[j] = t.Dependent[j], t.Dependent[i]
	t.Ind[i], t.Ind[j] = t.Ind[j], t.Ind[i]
}

// Less ...
func (t *SortSliceFloat64BySliceString) Less(i, j int) bool {
	if t.Ascending {
		return t.Ind.Less(i, j)
	}

	return t.Ind.Less(j, i)
}

// Dependents ...
func (t *SortSliceFloat64BySliceString) Dependents() interface{} {
	return t.Dependent
}

// Independent ...
func (t *SortSliceFloat64BySliceString) Independent() interface{} {
	return t.Ind
}

// Order ...
func (t *SortSliceFloat64BySliceString) Order() {
	sort.Sort(t)
}
