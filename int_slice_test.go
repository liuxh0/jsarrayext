package jsarrayext

import (
	"reflect"
	"testing"
)

func TestIntSliceEvery(t *testing.T) {
	// It should return true in case of empty slice.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 0))
		r := s.Every(func(int, int) bool {
			return false
		})

		if r != true {
			t.Error()
		}
	})

	// It should return true if function always returns true.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		r := s.Every(func(int, int) bool {
			return true
		})

		if r != true {
			t.Error()
		}
	})

	// It should return false if function returns false once.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		r := s.Every(func(element int, index int) bool {
			if index == 5 {
				return false
			}
			return true
		})

		if r != false {
			t.Error()
		}
	})

	// It should call the function with each element and index.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		nCalls := 0
		s.Every(func(element int, index int) bool {
			if element != s[nCalls] {
				t.Error()
			}

			if index != nCalls {
				t.Error()
			}

			nCalls++
			return true
		})

		if nCalls != len(s) {
			t.Error()
		}
	})

	// It should not call the function any more after false is returned.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))

		const falseIndex = 5
		nCalls := 0
		s.Every(func(element int, index int) bool {
			if index > falseIndex {
				t.Error()
			}

			nCalls++
			return index != falseIndex
		})

		if nCalls != falseIndex+1 {
			t.Error()
		}
	})
}

func TestIntSliceFill(t *testing.T) {
	// It should return the same slice.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 1))
		r := s.Fill(0, 0, 0)

		if &s[0] != &r[0] {
			t.Error()
		}
	})

	// It should fill with value.
	t.Run("", func(t *testing.T) {
		start, end := 1, 4
		s := IntSlice(make([]int, 5))
		r := s.Fill(1, start, end)

		r.ForEach(func(element int, index int) {
			if index >= start && index < end {
				if element != 1 {
					t.Error()
				}
			} else {
				if element != 0 {
					t.Error()
				}
			}
		})
	})
}

func TestIntSliceFilter(t *testing.T) {
	// It should return a slice with the right capacity.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10, 20))
		r := s.Filter(func(int, int) bool {
			return false
		})

		if len(r) != 0 || cap(r) != len(s) {
			t.Error()
		}
	})

	// It should call the function with each element and index.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		nCalls := 0
		s.Filter(func(element int, index int) bool {
			if element != s[nCalls] {
				t.Error()
			}

			if index != nCalls {
				t.Error()
			}

			nCalls++
			return false
		})

		if nCalls != len(s) {
			t.Error()
		}
	})

	// It should return a slice with filtered elements.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		r := s.Filter(func(element int, index int) bool {
			return element == 3 || element == 7
		})

		if len(r) != 2 {
			t.Error()
		}

		if r[0] != 3 || r[1] != 7 {
			t.Error()
		}
	})
}

func TestIntSliceFind(t *testing.T) {
	// It should call the function with each element and index.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		nCalls := 0
		s.Find(func(element int, index int) bool {
			if element != s[nCalls] {
				t.Error()
			}

			if index != nCalls {
				t.Error()
			}

			nCalls++
			return false
		})

		if nCalls != len(s) {
			t.Error()
		}
	})

	// It should not call the function any more after true is returned.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))

		const trueIndex = 5
		nCalls := 0
		s.Find(func(element int, index int) bool {
			if index > trueIndex {
				t.Error()
			}

			nCalls++
			return index == trueIndex
		})

		if nCalls != trueIndex+1 {
			t.Error()
		}
	})

	// It should return the element if found and return 0 if not.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		const findValue = 5
		r1 := s.Find(func(element int, index int) bool {
			return element == findValue
		})

		if r1 != findValue {
			t.Error()
		}

		const nonExistValue = 20
		r2 := s.Find(func(element int, index int) bool {
			return element == nonExistValue
		})

		if r2 != 0 {
			t.Error()
		}
	})
}

func TestIntSliceFindIndex(t *testing.T) {
	// It should call the function with each element and index.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		nCalls := 0
		s.FindIndex(func(element int, index int) bool {
			if element != s[nCalls] {
				t.Error()
			}

			if index != nCalls {
				t.Error()
			}

			nCalls++
			return false
		})

		if nCalls != len(s) {
			t.Error()
		}
	})

	// It should not call the function any more after true is returned.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))

		const trueIndex = 5
		nCalls := 0
		s.FindIndex(func(element int, index int) bool {
			if index > trueIndex {
				t.Error()
			}

			nCalls++
			return index == trueIndex
		})

		if nCalls != trueIndex+1 {
			t.Error()
		}
	})

	// It should return the index if found and return -1 if not.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		const findIndex = 5
		r1 := s.FindIndex(func(element int, index int) bool {
			return index == findIndex
		})

		if r1 != findIndex {
			t.Error()
		}

		const nonExistValue = 20
		r2 := s.FindIndex(func(element int, index int) bool {
			return element == nonExistValue
		})

		if r2 != -1 {
			t.Error()
		}
	})
}

func TestIntSliceForEach(t *testing.T) {
	// It should call the function with each element and index.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		nCalls := 0
		s.ForEach(func(element int, index int) {
			if element != s[nCalls] {
				t.Error()
			}

			if index != nCalls {
				t.Error()
			}

			nCalls++
		})

		if nCalls != len(s) {
			t.Error()
		}
	})
}

func TestIntSliceIncludes(t *testing.T) {
	// It should return true if the slice includes the value and return false on
	// the contrary.
	t.Run("", func(t *testing.T) {
		s := IntSlice([]int{1, 2, 3, 4, 5})
		rt := s.Includes(3)
		rf := s.Includes(6)

		if rt != true {
			t.Error()
		}

		if rf != false {
			t.Error()
		}
	})
}

func TestIntSliceIndexOf(t *testing.T) {
	// It should return the index if the value is found and return -1 if not.
	t.Run("", func(t *testing.T) {
		s := IntSlice([]int{1, 2, 3, 4, 5})
		r1 := s.IndexOf(3)
		r2 := s.IndexOf(6)

		if r1 != 2 {
			t.Error()
		}

		if r2 != -1 {
			t.Error()
		}
	})
}

func TestIntSliceLastIndexOf(t *testing.T) {
	// It should return the index if the value is found and return -1 if not.
	t.Run("", func(t *testing.T) {
		s := IntSlice([]int{1, 2, 3, 2, 1})
		r1 := s.LastIndexOf(2)
		r2 := s.LastIndexOf(4)

		if r1 != 3 {
			t.Error()
		}

		if r2 != -1 {
			t.Error()
		}
	})
}

func TestIntSliceMap(t *testing.T) {
	// It should return a slice with mapped values.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		r := s.Map(func(element int, index int) interface{} {
			return element * 10
		})

		for i := range r {
			if r[i] != s[i]*10 {
				t.Error()
			}
		}
	})
}

func TestIntSliceReduceRight(t *testing.T) {
	// It should return the correct result.
	t.Run("", func(t *testing.T) {
		s := IntSlice([]int{1, 2, 3})

		r := s.ReduceRight(func(previousValue interface{}, currentValue int, currentIndex int) interface{} {
			return previousValue.(int)*10 + currentValue
		}, 1)

		if r != 1321 {
			t.Error(r)
		}
	})
}

func TestIntSliceReverse(t *testing.T) {
	s := IntSlice([]int{1, 2, 3})
	r := s.Reverse()

	if reflect.DeepEqual([]int(r), []int{3, 2, 1}) == false {
		t.Error()
	}
}

func TestIntSliceSome(t *testing.T) {
	// It should return false in case of empty slice.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 0))
		r := s.Some(func(int, int) bool {
			return false
		})

		if r != false {
			t.Error()
		}
	})

	// It should return false if function always returns false.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		r := s.Some(func(int, int) bool {
			return false
		})

		if r != false {
			t.Error()
		}
	})

	// It should return true if function returns true once.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		r := s.Some(func(element int, index int) bool {
			if index == 5 {
				return true
			}
			return false
		})

		if r != true {
			t.Error()
		}
	})

	// It should call the function with each element and index.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))
		for i := range s {
			s[i] = i
		}

		nCalls := 0
		s.Some(func(element int, index int) bool {
			if element != s[nCalls] {
				t.Error()
			}

			if index != nCalls {
				t.Error()
			}

			nCalls++
			return false
		})

		if nCalls != len(s) {
			t.Error()
		}
	})

	// It should not call the function any more after true is returned.
	t.Run("", func(t *testing.T) {
		s := IntSlice(make([]int, 10))

		const falseIndex = 5
		nCalls := 0
		s.Some(func(element int, index int) bool {
			if index > falseIndex {
				t.Error()
			}

			nCalls++
			return index == falseIndex
		})

		if nCalls != falseIndex+1 {
			t.Error()
		}
	})
}

func TestIntSliceSort(t *testing.T) {
	s := IntSlice([]int{4, 3, 1, 5, 2})
	r := s.Sort(func(firstElement int, secondElement int) int {
		return firstElement - secondElement
	})

	if reflect.DeepEqual([]int(r), []int{1, 2, 3, 4, 5}) == false {
		t.Error()
	}
}
