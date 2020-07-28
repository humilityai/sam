package slices

// MapIntSliceInt ...
type MapIntSliceInt map[int]Int

// Add ...
func (m MapIntSliceInt) Add(i int) {
	_, ok := m[i]
	if !ok {
		m[i] = make(Int, 0)
	}
}

// MustAppend will create the integer key if it does
// not already exist, and otherwise append the value
// to it's slice.
func (m MapIntSliceInt) MustAppend(i, value int) {
	s, ok := m[i]
	if !ok {
		m[i] = Int{value}
	}

	s = append(s, value)
	m[i] = s
}

// AddWithLength ...
func (m MapIntSliceInt) AddWithLength(i, length int) {
	_, ok := m[i]
	if !ok {
		m[i] = make(Int, length, length)
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
