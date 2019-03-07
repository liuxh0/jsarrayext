package jsarrayext

import (
	"reflect"
	"sort"
)

func concat(slices ...interface{}) interface{} {
	newSliceLen := 0
	for _, v := range slices {
		newSliceLen += reflect.ValueOf(v).Len()
	}

	newSlice := make([]interface{}, newSliceLen)
	index := 0
	for _, slice := range slices {
		sliceValue := reflect.ValueOf(slice)
		sliceLen := sliceValue.Len()
		for i := 0; i < sliceLen; i++ {
			newSlice[index] = sliceValue.Index(i).Interface()
			index++
		}
	}

	return newSlice
}

func every(
	slice interface{},
	fn func(element interface{}, index int) bool,
) bool {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index) == false {
			return false
		}
	}

	return true
}

func fill(
	slice interface{},
	value interface{},
	start int,
	end int,
) interface{} {
	val := func() reflect.Value {
		if value == nil {
			return reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
		}
		return reflect.ValueOf(value)
	}()

	for index := start; index < end; index++ {
		reflect.ValueOf(slice).Index(index).Set(val)
	}

	return slice
}

func filter(
	slice interface{},
	fn func(element interface{}, index int) bool,
) interface{} {
	filtered := reflect.MakeSlice(reflect.TypeOf(slice), 0, reflect.ValueOf(slice).Len())

	forEach(slice, func(element interface{}, index int) {
		if fn(element, index) == true {
			if element == nil {
				// reflect.Append(filtered, nil) won't work.
				// Instead create a zero value of the type and append it.
				// See also https://groups.google.com/forum/#!topic/golang-nuts/Txje1_UfaMQ
				filtered = reflect.Append(filtered, reflect.Zero(reflect.TypeOf(slice).Elem()))
			} else {
				filtered = reflect.Append(filtered, reflect.ValueOf(element))
			}
		}
	})

	return filtered.Interface()
}

func find(
	slice interface{},
	fn func(element interface{}, index int) bool,
) interface{} {
	if index := findIndex(slice, fn); index != -1 {
		element := reflect.ValueOf(slice).Index(index).Interface()
		return element
	}

	return nil
}

func findIndex(
	slice interface{},
	fn func(element interface{}, index int) bool,
) int {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index) == true {
			return index
		}
	}

	return -1
}

func forEach(
	slice interface{},
	fn func(element interface{}, index int),
) {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		fn(element, index)
	}
}

func includes(
	slice interface{},
	value interface{},
) bool {
	return some(slice, func(element interface{}, index int) bool {
		return reflect.DeepEqual(element, value)
	})
}

func indexOf(
	slice interface{},
	value interface{},
) int {
	return findIndex(slice, func(element interface{}, index int) bool {
		return reflect.DeepEqual(element, value)
	})
}

func lastIndexOf(
	slice interface{},
	value interface{},
) int {
	sliceValue := reflect.ValueOf(slice)
	len := sliceValue.Len()

	for index := len - 1; index >= 0; index-- {
		element := sliceValue.Index(index).Interface()
		if reflect.DeepEqual(element, value) {
			return index
		}
	}

	return -1
}

func mapToInterfaceSlice(
	slice interface{},
	fn func(element interface{}, index int) interface{},
) []interface{} {
	mapped := make([]interface{}, reflect.ValueOf(slice).Len())

	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		mapped[index] = fn(element, index)
	}

	return mapped
}

func reduce(
	slice interface{},
	fn func(previousValue interface{}, currentValue interface{}, currentIndex int) interface{},
	initialValue interface{},
) interface{} {
	previousValue := initialValue
	sliceValue := reflect.ValueOf(slice)
	sliceLen := sliceValue.Len()

	for index := 0; index < sliceLen; index++ {
		currentValue := sliceValue.Index(index).Interface()
		previousValue = fn(previousValue, currentValue, index)
	}

	return previousValue
}

func reduceRight(
	slice interface{},
	fn func(previousValue interface{}, currentValue interface{}, currentIndex int) interface{},
	initialValue interface{},
) interface{} {
	previousValue := initialValue
	sliceValue := reflect.ValueOf(slice)
	sliceLen := sliceValue.Len()

	for index := sliceLen - 1; index >= 0; index-- {
		currentValue := sliceValue.Index(index).Interface()
		previousValue = fn(previousValue, currentValue, index)
	}

	return previousValue
}

func reverse(
	slice interface{},
) interface{} {
	sliceLen := reflect.ValueOf(slice).Len()
	swapper := reflect.Swapper(slice)

	for index := 0; index < (sliceLen / 2); index++ {
		reversedIndex := sliceLen - index - 1
		if index != reversedIndex {
			swapper(index, reversedIndex)
		}
	}

	return slice
}

func some(
	slice interface{},
	fn func(element interface{}, index int) bool,
) bool {
	for index := 0; index < reflect.ValueOf(slice).Len(); index++ {
		element := reflect.ValueOf(slice).Index(index).Interface()
		if fn(element, index) == true {
			return true
		}
	}

	return false
}

func sortSlice(
	slice interface{},
	fn func(firstElement interface{}, secondElement interface{}) int,
) interface{} {
	sliceValue := reflect.ValueOf(slice)

	sort.Slice(slice, func(i, j int) bool {
		firstElement := sliceValue.Index(i).Interface()
		secondElement := sliceValue.Index(j).Interface()
		return fn(firstElement, secondElement) < 0
	})

	return slice
}
