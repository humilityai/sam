package sam

import "math"

// MapStringInt can be built by simply calling
// make(MapStringInt).
type MapStringInt map[string]int

// Increment will add 1 to the integer value found
// at the provided key.
func (m MapStringInt) Increment(key string) {
	v, ok := m[key]
	if !ok {
		m[key] = 1
	} else {
		m[key] = v + 1
	}
}

// Values will return the slice of all values
// found in the map.
func (m MapStringInt) Values() (values SliceInt) {
	for _, v := range m {
		values = append(values, v)
	}

	return
}

// Keys will return the slice of all keys
// found in the map.
func (m MapStringInt) Keys() (keys SliceString) {
	for k := range m {
		keys = append(keys, k)
	}

	return
}

// MaxValue will return the largest value and it's key.
// The key is the first result returned. And the value
// is the second result returned.
func (m MapStringInt) MaxValue() (key string, value int) {
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
func (m MapStringInt) MinValue() (key string, value int) {
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
