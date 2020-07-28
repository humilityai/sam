package sam

import "math"

// Int ...
type Int []int

// Product ...
func (s Int) Product() int {
	var product int
	for i, v := range s {
		if i == 0 {
			product = v
			continue
		}
		product *= v
	}

	return product
}

// Sum ...
func (s Int) Sum() int {
	var sum int

	for i, v := range s {
		if i == 0 {
			sum = v
			continue
		}
		sum += v
	}

	return sum
}

// Min ...
func (s Int) Min() (index, value int) {
	value = math.MaxInt64
	for i, v := range s {
		if v < value {
			index = i
			value = v
		}
	}

	return
}

// Max ...
func (s Int) Max() (index int) {
	max := math.MinInt64
	for i, v := range s {
		if v > max {
			v = max
			index = i
		}
	}

	return
}

// Contains ...
func (s Int) Contains(i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

// IncrementPosition ...
func (s Int) IncrementPosition(index int) {
	if index > len(s)-1 || index < 0 {
		return
	}

	s[index]++
}
