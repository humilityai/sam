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

// MapIntSliceInt uses integer values as a key
// and a slice of integers as values.
type MapIntSliceInt map[int]SliceInt

// Add ...
func (m MapIntSliceInt) Add(i int) {
	_, ok := m[i]
	if !ok {
		m[i] = make(SliceInt, 0)
	}
}

// MustAppend will create the integer key if it does
// not already exist, and otherwise append the value
// to it's slice.
func (m MapIntSliceInt) MustAppend(i, value int) {
	s, ok := m[i]
	if !ok {
		m[i] = SliceInt{value}
	}

	s = append(s, value)
	m[i] = s
}

// AddWithLength ...
func (m MapIntSliceInt) AddWithLength(i, length int) {
	_, ok := m[i]
	if !ok {
		m[i] = make(SliceInt, length, length)
	}
}

// AppendSlice ...
func (m MapIntSliceInt) AppendSlice(i, value int) {
	s, ok := m[i]
	if !ok {
		return
	}

	s = append(s, value)
	m[i] = s
}

// InsertSliceValue ...
func (m MapIntSliceInt) InsertSliceValue(i, value, position int) {
	s, ok := m[i]
	if !ok {
		return
	}

	s[position] = value
	m[i] = s
}

// IncrementSliceValue ...
func (m MapIntSliceInt) IncrementSliceValue(i, value, position int) {
	s, ok := m[i]
	if !ok {
		return
	}
	s.IncrementPosition(position)

	m[i] = s
}
