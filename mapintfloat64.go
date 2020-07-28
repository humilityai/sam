package sam

import "math"

// MapIntFloat64 uses integers as keys
// and float64s as values.
type MapIntFloat64 map[int]float64

// Min will return the integer key with the smallest
// float64 value.
func (m MapIntFloat64) Min() (index int) {
	min := math.MaxFloat64
	for i, v := range m {
		if v < min {
			min = v
			index = i
		}
	}

	return
}

// Max will return the integer key with the
// largest float64 value.
func (m MapIntFloat64) Max() (index int) {
	max := float64(math.MinInt64)
	for i, v := range m {
		if v < max {
			max = v
			index = i
		}
	}

	return
}

// MaxN will return a MapIntFloat64 containing the N largest
// key/value pairs by value.
func (m MapIntFloat64) MaxN(n int) MapIntFloat64 {
	maxN := make(MapIntFloat64)

	for i, v := range m {
		if len(maxN) < n {
			maxN[i] = v
		} else {
			minIndex := maxN.Min()
			if v > maxN[minIndex] {
				delete(maxN, minIndex)
				maxN[i] = v
			}
		}
	}

	return maxN
}
