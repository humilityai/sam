package sam

import "math"

// MapStringFloat64 can be built by simply calling
// make(MapStringFloat64).
type MapStringFloat64 map[string]float64

// Increment will add 1 to the float64 value found
// at the provided key.
func (m MapStringFloat64) Increment(key string) {
	m[key]++
}

// Contains will return whether or not a key exists
// in the map.
func (m MapStringFloat64) Contains(s string) bool {
	_, ok := m[s]
	return ok
}

// Values will return the slice of all values
// found in the map.
func (m MapStringFloat64) Values() (values SliceFloat64) {
	for _, v := range m {
		values = append(values, v)
	}

	return
}

// Keys will return the slice of all keys
// found in the map.
func (m MapStringFloat64) Keys() (keys SliceString) {
	for k := range m {
		keys = append(keys, k)
	}

	return
}

// KeysAndValues will return two ordered sets of keys and values that are aligned
// by slice index.
func (m MapStringFloat64) KeysAndValues() (keys SliceString, values SliceFloat64) {
	keys = make(SliceString, 0, len(m))
	values = make(SliceFloat64, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	return
}

// MaxValue will return the largest value and it's key.
// The key is the first result returned. And the value
// is the second result returned.
func (m MapStringFloat64) MaxValue() (key string, value float64) {
	max := math.SmallestNonzeroFloat64
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
func (m MapStringFloat64) MinValue() (key string, value float64) {
	min := math.MaxFloat64
	for k, v := range m {
		if v < min {
			min = v
			key = k
			value = v
		}
	}

	return
}
