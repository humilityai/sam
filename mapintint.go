package slices

import "math"

// MapIntInt can be built by simply calling
// make(MapIntInt).
type MapIntInt map[int]int

// Increment will add 1 to the integer value found
// at the provided key.
func (m MapIntInt) Increment(key int) {
	v, ok := m[key]
	if !ok {
		m[key] = 1
	} else {
		m[key] = v + 1
	}
}

// Values will return the slice of all values
// found in the map.
func (m MapIntInt) Values() (values []int) {
	for _, v := range m {
		values = append(values, v)
	}

	return
}

// Keys will return the slice of all keys
// found in the map.
func (m MapIntInt) Keys() (keys []int) {
	for k := range m {
		keys = append(keys, k)
	}

	return
}

// MaxValue will return he largest value and it's key.
// The key is the first result returned. And the value
// is the second result returned.
func (m MapIntInt) MaxValue() (key int, value int) {
	max := math.MinInt64
	for k, v := range m {
		if v > max {
			max = v
			key = k
			value = v
		}
	}

	return
}

// MinValue will return the smallest value and it's key.
// The key is the first result returned. And the value
// is the second result returned.
func (m MapIntInt) MinValue() (key int, value int) {
	min := math.MaxInt64
	for k, v := range m {
		if v < min {
			min = v
			key = k
			value = v
		}
	}

	return
}

// Contains will check if the map contains a key.
func (m MapIntInt) Contains(key int) bool {
	for k := range m {
		if k == key {
			return true
		}
	}

	return false
}
