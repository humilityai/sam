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
