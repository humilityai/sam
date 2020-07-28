package slices

import "math"

// MapFloat64Int ...
type MapFloat64Int map[float64]int

// Max ...
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

// Increment ...
func (m MapFloat64Int) Increment(f float64) {
	v, ok := m[f]
	if !ok {
		m[f] = 1
		return
	}
	v++

	m[f] = v
}

// AverageCount ...
func (m MapFloat64Int) AverageCount() int {
	var avg int

	for _, count := range m {
		avg += count
	}
	avg = avg / len(m)

	return avg
}
