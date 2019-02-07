package jsarrayext

// Slice ...
type Slice []interface{}

// Every ...
func (s Slice) Every(fn func(element interface{}, index int) bool) bool {
	return every(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}

// Fill fills elements from a start index to an end index (not included) with a
// static value.
func (s Slice) Fill(value interface{}, start int, end int) Slice {
	return fill(s, value, start, end).(Slice)
}

// Filter ...
func (s Slice) Filter(fn func(element interface{}, index int) bool) Slice {
	return filter(s, func(element interface{}, index int) bool {
		return fn(element, index)
	}).(Slice)
}

// Find ...
func (s Slice) Find(fn func(element interface{}, index int) bool) interface{} {
	return find(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}

// FindIndex ...
func (s Slice) FindIndex(fn func(element interface{}, index int) bool) int {
	return findIndex(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}

// ForEach ...
func (s Slice) ForEach(fn func(element interface{}, index int)) {
	forEach(s, func(element interface{}, index int) {
		fn(element, index)
	})
}

// Includes determines whether a slice includes a certain value. Equality is
// determined by reflect.DeepEqual().
func (s Slice) Includes(value interface{}) bool {
	return includes(s, value)
}

// IndexOf returns the first index at which a given value can be found in the
// slice, or -1 if it is not present. Equality is determined by
// reflect.DeepEqual().
func (s Slice) IndexOf(value interface{}) int {
	return indexOf(s, value)
}

// LastIndexOf returns the last index at which a given value can be found in the
// slice, or -1 if it is not present. Equality is determined by
// reflect.DeepEqual().
func (s Slice) LastIndexOf(value interface{}) int {
	return lastIndexOf(s, value)
}

// Map ...
func (s Slice) Map(fn func(element interface{}, index int) interface{}) Slice {
	return mapToInterfaceSlice(s, func(element interface{}, index int) interface{} {
		return fn(element, index)
	})
}

// Reduce executes the reducer function on each element of the array and returns
// a single output value.
func (s Slice) Reduce(
	fn func(previousValue interface{}, currentValue interface{}, currentIndex int) interface{},
	initialValue interface{},
) interface{} {
	return reduce(s, fn, initialValue)
}

// Some ...
func (s Slice) Some(fn func(element interface{}, index int) bool) bool {
	return some(s, func(element interface{}, index int) bool {
		return fn(element, index)
	})
}
