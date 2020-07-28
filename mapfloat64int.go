package sam

import "math"

// MapFloat64Int uses float64s as keys
// and integers as values.
type MapFloat64Int map[float64]int

// Max will return the key with the largest
// integer value.
func (m MapFloat64Int) Max() float64 {
	max := math.MinInt64
	var returnValue float64
	for value, count := range m {
		if count > max {
			returnValue = value
			max = count
		}
	}

	return returnValue
}

// Increment will add 1 to the integer value
// of the provided float64 key.
func (m MapFloat64Int) Increment(f float64) {
	v, ok := m[f]
	if !ok {
		m[f] = 1
		return
	}
	v++

	m[f] = v
}

// AverageCount will iterate over the map and
// return the average integer value.
func (m MapFloat64Int) AverageCount() int {
	var avg int

	for _, count := range m {
		avg += count
	}
	avg = avg / len(m)

	return avg
}
