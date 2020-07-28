package sam

import (
	"math"
	"math/rand"
	"sort"

	. "github.com/humilityai/util"
)

// SliceFloat64 is the primary data structure for holding numerical
// data.
type SliceFloat64 []float64

// MultiplyBy will mutliply all values in the slice
// by the value provided.
func (s SliceFloat64) MultiplyBy(x float64) {
	for i, value := range s {
		s[i] = value * x
	}
}

// Equal ...
func (s SliceFloat64) Equal(element interface{}) bool {
	input, ok := element.(SliceFloat64)
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

// Min ...
func (s SliceFloat64) Min() float64 {
	if len(s) == 0 {
		return 0
	}
	min := math.MaxFloat64

	for _, value := range s {
		if value < min {
			min = value
		}
	}

	return min
}

// MinIndex ...
func (s SliceFloat64) MinIndex() (int, float64) {
	min := math.MaxFloat64
	var index int

	for i, value := range s {
		if value < min {
			min = value
			index = i
		}
	}

	return index, min
}

// AverageDeviation ...
func (s SliceFloat64) AverageDeviation() float64 {
	avg := s.Avg()
	var totalDev float64
	for _, v := range s {
		dev := math.Abs(v - avg)
		totalDev += dev
	}

	return totalDev / float64(len(s))
}

// IsNum will convert any non-numerical values
// in the slice into a numerical value (0) and return
// a new slice.
func (s SliceFloat64) IsNum() SliceFloat64 {
	newS := make(SliceFloat64, len(s), len(s))
	for i, v := range s {
		newS[i] = IsNum(v)
	}

	return newS
}

// Median ...
func (s SliceFloat64) Median() (index int, value float64) {
	if len(s) == 0 {
		return 0, 0
	}
	sort.Sort(s)

	index = len(s) / 2
	value = s[index]

	return
}

// Sum will return the sum of all the values
// in the SliceFloat64.
func (s SliceFloat64) Sum() (sum float64) {
	for _, v := range s {
		sum += v
	}

	return
}

// Max will return the largest value in the slice.
func (s SliceFloat64) Max() float64 {
	if len(s) == 0 {
		return 0
	}
	max := math.SmallestNonzeroFloat64

	for _, value := range s {
		if value > max {
			max = value
		}
	}

	return max
}

// MaxN will return a []float64 of the largest-N
// values in the slice.
func (s SliceFloat64) MaxN(n int) SliceFloat64 {
	if n <= 0 {
		return make(SliceFloat64, 0)
	}

	var maxN SliceFloat64
	for _, value := range s {
		if len(maxN) < n {
			maxN = append(maxN, value)
		} else {
			minIndex, minValue := maxN.MinIndex()
			if value > minValue {
				maxN[minIndex] = value
			}
		}
	}

	return maxN
}

// MaxNWithIndex will return a map[int]float64 of the largest-N values
// and the index that they reside at: map[index]value
func (s SliceFloat64) MaxNWithIndex(n int) MapIntFloat64 {
	if n <= 0 {
		return make(MapIntFloat64)
	}

	var maxN MapIntFloat64
	for i, v := range s {
		if len(maxN) < n {
			maxN[i] = v
		} else {
			minIndex := maxN.Min()
			if v > maxN[minIndex] {
				maxN[minIndex] = v
			}
		}
	}

	return maxN
}

// Avg will return the mean of the values
// in the slice.
func (s SliceFloat64) Avg() float64 {
	var total float64
	for _, value := range s {
		total += value
	}

	avg := total / float64(len(s))

	return avg
}

// MajorityZero will return true iff the number of
// zeros in the data is greater than half of the values.
func (s SliceFloat64) MajorityZero() bool {
	var totalZeros int

	for _, value := range s {
		if value == 0 {
			totalZeros++
		}
	}

	percentage := float64(totalZeros) / float64(len(s))
	if percentage > 0.5 {
		return true
	}

	return false
}

// Bounds will return the maximum and minimum values in the slice.
func (s SliceFloat64) Bounds() (min float64, max float64) {
	max = math.SmallestNonzeroFloat64
	min = math.MaxFloat64

	for _, value := range s {

		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return
}

// NonZeroValues ...
func (s SliceFloat64) NonZeroValues() SliceFloat64 {
	var nonZeroValues SliceFloat64
	for _, value := range s {
		if value != float64(0) {
			nonZeroValues = append(nonZeroValues, value)
		}
	}

	return nonZeroValues
}

// NonZeroBounds will return the non-zero values which are the minimum
// and maximum values in the slice. But, if slice is zeroed, then
// it will return 0 for both min and max.
func (s SliceFloat64) NonZeroBounds() (min float64, max float64) {
	max = math.SmallestNonzeroFloat64
	min = math.MaxFloat64

	if s.IsZeroed() {
		min = 0
		max = 0
		return
	}

	for _, value := range s {
		if value > max && value != 0 {
			max = value
		}

		if value < min && value != 0 {
			min = value
		}
	}

	return
}

// Range ...
func (s SliceFloat64) Range() float64 {
	max := math.SmallestNonzeroFloat64
	min := math.MaxFloat64
	for _, value := range s {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return math.Abs(max - min)
}

// Mode ...
func (s SliceFloat64) Mode() float64 {
	counts := make(MapFloat64Int)
	for _, value := range s {
		v, ok := counts[value]
		if !ok {
			counts[value] = 1
			continue
		}

		v++
		counts[value] = v
	}

	return counts.Max()
}

// Counts ...
func (s SliceFloat64) Counts() MapFloat64Int {
	counts := make(MapFloat64Int)
	for _, v := range s {
		counts.Increment(v)
	}

	return counts
}

// Contains is technically equivalent to the EqualTo method.
// It simply resolves faster iff the slice actually contains
// the supplied argument. Use this method for better best-case performance and
// a guaranteed smaller memory footprint.
func (s SliceFloat64) Contains(f float64) bool {
	for _, value := range s {
		if value == f {
			return true
		}
	}

	return false
}

// GreaterThan will return the indices, total count, and percentage of slice of the float64 numbers
// which are greater than the supplied float64 argument.
func (s SliceFloat64) GreaterThan(value float64) (indices []int, count int, percentage float64) {
	for i, v := range s {
		if v > value {
			indices = append(indices, i)
			count++
		}
	}

	percentage = float64(count) / float64(len(s))

	return
}

// LessThan will return the indices, total count, and percentage of slice of the float64 numbers
// which are less than the supplied float64 argument.
func (s SliceFloat64) LessThan(value float64) (indices []int, count int, percentage float64) {
	for i, v := range s {
		if v < value {
			indices = append(indices, i)
			count++
		}
	}

	percentage = float64(count) / float64(len(s))

	return
}

// EqualTo will return the indices, total count, and percentage of slice of the float64 numbers
// which are equal to the supplied float64 argument.
func (s SliceFloat64) EqualTo(value float64) (indices []int, count int, percentage float64) {
	for i, v := range s {
		if v == value {
			indices = append(indices, i)
			count++
		}
	}

	percentage = float64(count) / float64(len(s))

	return
}

// IsZeroed ...
func (s SliceFloat64) IsZeroed() bool {
	for _, value := range s {
		if value != 0 {
			return false
		}
	}

	return true
}

// StrictlyPositive ...
func (s SliceFloat64) StrictlyPositive() bool {
	for _, value := range s {
		if value < 0 {
			return false
		}
	}

	return true
}

// StrictlyNegative ...
func (s SliceFloat64) StrictlyNegative() bool {
	for _, value := range s {
		if value < 0 {
			return false
		}
	}

	return true
}

// RescaleValues will rescale the values to either [0,1]
// or [-1,1] range. It returns the min and max values found
// in the scalar slice.
func (s SliceFloat64) RescaleValues() (min, max float64) {
	min = s.Min()
	max = s.Max()
	// avg := stat.Mean(values, nil)

	scaledValues := make(SliceFloat64, len(s), len(s))
	for index, value := range s {
		scaledValue := (value - min) / (max - min)
		scaledValues[index] = scaledValue
	}

	s = scaledValues
	return
}

// ShiftLogScaleValues ...
func (s SliceFloat64) ShiftLogScaleValues() (shift float64) {
	minValue := s.Min()
	if minValue <= 0 {
		shift = (0 - minValue) + 1
	}

	newData := make(SliceFloat64, len(s), len(s))
	for index, value := range s {
		shiftedValue := value + shift
		scaledValue := math.Log(shiftedValue)
		newData[index] = scaledValue
	}

	s = newData

	return
}

// MeanCrossingRate ...
func (s SliceFloat64) MeanCrossingRate() float64 {
	mean := s.Avg()

	var crossings int
	for i := range s {
		if i < len(s)-1 {
			if s[i] < mean && s[i+1] > mean {
				crossings++
			} else if s[i] > mean && s[i+1] < mean {
				crossings++
			}
		}
	}

	// crossing rate
	rate := float64(crossings) / float64(len(s))

	return rate
}

// ZeroingRate should be called instead of ZeroCrossingRate
// or ZeroHittingRate, as it will automatically choose the best
// function to use and return the result.
func (s SliceFloat64) ZeroingRate() (rate float64) {
	if s.StrictlyPositive() || s.StrictlyNegative() {
		rate = s.ZeroHittingRate()
		return
	}

	s.ZeroCrossingRate()
	return
}

// ZeroCrossingRate counts the number of times the slice
// is above 0 and then below zero (in index order), or vice versa.
// Consider calling ZeroingRate method instead.
func (s SliceFloat64) ZeroCrossingRate() float64 {
	if s.StrictlyPositive() || s.StrictlyNegative() {
		return 0
	}

	var crossings int
	for i := range s {
		if i < len(s)-1 {
			if s[i] < 0 && s[i+1] > 0 {
				crossings++
			} else if s[i] > 0 && s[i+1] < 0 {
				crossings++
			}
		}
	}

	// crossing rate
	rate := float64(crossings) / float64(len(s))

	return rate
}

// ZeroHittingRate is used for when the data is either strictly positive
// or strictly negative.
// Consider calling ZeroingRate method instead.
func (s SliceFloat64) ZeroHittingRate() float64 {
	var hits int
	if s.StrictlyPositive() || s.StrictlyNegative() {
		for _, value := range s {
			if value == 0 {
				hits++
			}
		}
	}

	// zero-hit rate
	rate := float64(hits) / float64(len(s))

	return rate
}

// Peaks ...
func (s SliceFloat64) Peaks() (peaks map[int]float64) {
	peaks = make(map[int]float64)

	for i, v := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] > s[i-1] && s[i] > s[i+1] {
				peaks[i] = v
			}
		}
	}

	return
}

// Vallies ...
func (s SliceFloat64) Vallies() (vallies map[int]float64) {
	vallies = make(map[int]float64)

	for i, v := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] < s[i-1] && s[i] < s[i+1] {
				vallies[i] = v
			}
		}
	}

	return
}

// Extrema ...
func (s SliceFloat64) Extrema() (extrema map[int]float64) {
	extrema = make(map[int]float64)

	for i, v := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] < s[i-1] && s[i] < s[i+1] {
				// peaks
				extrema[i] = v
			} else if s[i] > s[i-1] && s[i] > s[i+1] {
				// valleys
				extrema[i] = v
			}
		}
	}

	return
}

// PeakCount ...
func (s SliceFloat64) PeakCount() (peaks int) {
	for i := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] > s[i-1] && s[i] > s[i+1] {
				peaks++
			}
		}
	}

	return
}

// ValleyCount ...
func (s SliceFloat64) ValleyCount() (vallies int) {
	for i := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] < s[i-1] && s[i] < s[i+1] {
				vallies++
			}
		}
	}

	return
}

// ValleyDistances ...
func (s SliceFloat64) ValleyDistances() (distances SliceFloat64) {
	var indices []int

	for i := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] < s[i-1] && s[i] < s[i+1] {
				indices = append(indices, i)
			}
		}
	}

	for j, index := range indices {
		if j > 0 {
			distance := float64(index - indices[j-1])
			distances = append(distances, distance)
		}
	}

	return
}

// PeakDistances ...
func (s SliceFloat64) PeakDistances() (distances SliceFloat64) {
	var indices []int

	for i := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] > s[i-1] && s[i] > s[i+1] {
				indices = append(indices, i)
			}
		}
	}

	for j, index := range indices {
		if j > 0 {
			distance := float64(index - indices[j-1])
			distances = append(distances, distance)
		}
	}

	return
}

// ExtremaDistances ...
func (s SliceFloat64) ExtremaDistances() (distances SliceFloat64) {
	var indices []int

	for i := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] > s[i-1] && s[i] > s[i+1] {
				indices = append(indices, i)
			} else if s[i] < s[i-1] && s[i] < s[i+1] {
				indices = append(indices, i)
			}
		}
	}

	for j, index := range indices {
		if j > 0 {
			distance := float64(index - indices[j-1])
			distances = append(distances, distance)
		}
	}

	return
}

// EqualToSlice will check if two SliceFloat64s are the same length
// and have the same values in the same order.
func (s SliceFloat64) EqualToSlice(s2 SliceFloat64) bool {
	if len(s) != len(s2) {
		return false
	}

	for i, v := range s {
		if s2[i] != v {
			return false
		}
	}

	return true
}

// SimilarTo will check if two SliceFloat64s are the same length
// and contain the same elements (but not necessarily in the same order).
func (s SliceFloat64) SimilarTo(s2 SliceFloat64) bool {
	if len(s) != len(s2) {
		return false
	}

	for _, v := range s {
		if !s2.Contains(v) {
			return false
		}
	}

	return true
}

// SamplePercentage will return a random sampling of the slice based on the percentage value.
// Percentage values must be between 0 and 1.
func (s SliceFloat64) SamplePercentage(percentage float64) SliceFloat64 {
	var sample SliceFloat64

	if percentage < 0 || percentage > 1 {
		return sample
	}

	percentage = percentage * 100
	for _, value := range s {
		if float64(rand.Intn(100)) < percentage {
			sample = append(sample, value)
		}
	}

	return sample
}

// SampleAmount will return a random sample of count `amount`.
func (s SliceFloat64) SampleAmount(amount int) SliceFloat64 {
	var sample SliceFloat64
	if amount < 0 {
		return sample
	}

	if amount > len(s) {
		return s
	}

	percentage := (float64(amount) / float64(len(s))) * 100
	for _, value := range s {
		if float64(rand.Intn(100)) < percentage {
			sample = append(sample, value)
		}
	}

	return sample
}

// Sequences will return all sub-arrays of the total array of the length
// argument, and seperated by the stride argument
func (s SliceFloat64) Sequences(length, stride int) [][]float64 {
	sequences := make([][]float64, 0)

	start := length + stride
	for i := range s {
		if i+length < len(s)-1 && i >= start {
			sequences = append(sequences, s[i:i+length])
			start = i + stride
		}
	}

	return sequences
}

// SequenceSample will return a randomly sampled sequence from
// the slice.
func (s SliceFloat64) SequenceSample(length int) SliceFloat64 {
	if length < 0 {
		return SliceFloat64{}
	}

	if length > len(s) {
		return s
	}

	i := rand.Intn(len(s) - length)
	return s[i : i+length]
}

// ConsecutiveSequences returns all sequences collected where stride = length
func (s SliceFloat64) ConsecutiveSequences(length int) [][]float64 {
	return s.Sequences(length, length)
}

// OverlappingSequences are sequences where stride = 1
func (s SliceFloat64) OverlappingSequences(length int) [][]float64 {
	return s.Sequences(length, 1)
}

// Len ...
func (s SliceFloat64) Len() int {
	return len(s)
}

// Swap ...
func (s SliceFloat64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less ...
func (s SliceFloat64) Less(i, j int) bool {
	return s[i] < s[j]
}
