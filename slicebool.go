package sam

// SliceBool is a slice/array of boolean values.
type SliceBool []bool

// Equal will check if SliceBool is equal to provided
// object.
func (s SliceBool) Equal(element interface{}) bool {
	input, ok := element.(SliceBool)
	if !ok {
		return false
	}

	if len(input) != len(s) {
		return false
	}

	for i, v := range s {
		if v != input[i] {
			return false
		}
	}

	return true
}

// At is just a function version of what can be done
// more easily with `[index]`.
func (s SliceBool) At(index int) bool {
	return s[index]
}

// FalseIndices will return the list of all indices
// in the slice that have a false value.
func (s SliceBool) FalseIndices() SliceInt {
	var indices SliceInt
	for i, v := range s {
		if !v {
			indices = append(indices, i)
		}
	}

	return indices
}

// FalseCount will return the total number of
// false values in the boolean slice.
func (s SliceBool) FalseCount() (count int) {
	for _, v := range s {
		if !v {
			count++
		}
	}

	return
}

// FalsePercentage will return the percentage of values
// that are false.
func (s SliceBool) FalsePercentage() float64 {
	return float64(s.FalseCount()) / float64(len(s))
}

// TrueIndices will return the list of all indices
// in the slice that have a true value.
func (s SliceBool) TrueIndices() SliceInt {
	var indices SliceInt
	for i, v := range s {
		if v {
			indices = append(indices, i)
		}
	}

	return indices
}

// TrueCount will return the total number of
// true values in the boolean slice.
func (s SliceBool) TrueCount() (count int) {
	for _, v := range s {
		if v {
			count++
		}
	}

	return
}

// TruePercentage will return the percentage of values
// that are true.
func (s SliceBool) TruePercentage() float64 {
	return float64(s.TrueCount()) / float64(len(s))
}
