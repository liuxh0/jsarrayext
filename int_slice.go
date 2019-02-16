package jsarrayext

// IntSlice ...
type IntSlice []int

// Every ...
func (s IntSlice) Every(fn func(element int, index int) bool) bool {
	return every(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})
}

// Fill fills elements from a start index to an end index (not included) with a
// static value.
func (s IntSlice) Fill(value int, start int, end int) IntSlice {
	return fill(s, value, start, end).(IntSlice)
}

// Filter ...
func (s IntSlice) Filter(fn func(element int, index int) bool) IntSlice {
	return filter(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	}).(IntSlice)
}

// Find ...
func (s IntSlice) Find(fn func(element int, index int) bool) int {
	r := find(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})

	if r == nil {
		return 0
	}
	return r.(int)
}

// FindIndex ...
func (s IntSlice) FindIndex(fn func(element int, index int) bool) int {
	return findIndex(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})
}

// ForEach ...
func (s IntSlice) ForEach(fn func(element int, index int)) {
	forEach(s, func(element interface{}, index int) {
		fn(element.(int), index)
	})
}

// Includes determines whether a slice includes a certain value.
func (s IntSlice) Includes(value int) bool {
	return includes(s, value)
}

// IndexOf returns the first index at which a given value can be found in the
// slice, or -1 if it is not present.
func (s IntSlice) IndexOf(value int) int {
	return indexOf(s, value)
}

// LastIndexOf returns the last index at which a given value can be found in the
// slice, or -1 if it is not present.
func (s IntSlice) LastIndexOf(value int) int {
	return lastIndexOf(s, value)
}

// Map ...
func (s IntSlice) Map(fn func(element int, index int) interface{}) Slice {
	return mapToInterfaceSlice(s, func(element interface{}, index int) interface{} {
		return fn(element.(int), index)
	})
}

// Reduce executes the reducer function on each element of the array and returns
// a single output value.
func (s IntSlice) Reduce(
	fn func(previousValue interface{}, currentValue int, currentIndex int) interface{},
	initialValue interface{},
) interface{} {
	return reduce(s, func(previousValue interface{}, currentValue interface{}, currentIndex int) interface{} {
		return fn(previousValue, currentIndex, currentIndex)
	}, initialValue)
}

// ReduceRight executes the reducer function on each element of the array from
// right to left and returns a single output value.
func (s IntSlice) ReduceRight(
	fn func(previousValue interface{}, currentValue int, currentIndex int) interface{},
	initialValue interface{},
) interface{} {
	return reduceRight(s, func(previousValue interface{}, currentValue interface{}, currentIndex int) interface{} {
		return fn(previousValue, currentValue.(int), currentIndex)
	}, initialValue)
}

// Some ...
func (s IntSlice) Some(fn func(element int, index int) bool) bool {
	return some(s, func(element interface{}, index int) bool {
		return fn(element.(int), index)
	})
}
