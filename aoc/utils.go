package aoc

import (
	"strconv"
)

func ToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

type Collector[T any] struct {
	Values         [][]T
	perArrayAmount int
}

func NewCollector[T any](perArrayAmount int) *Collector[T] {
	// create a new TCollection pointer
	tc := &Collector[T]{}
	// initialize the Values slice with a capacity of 3
	tc.Values = make([][]T, 0, 3)
	tc.perArrayAmount = perArrayAmount
	return tc
}

func (tc *Collector[T]) Accept(t T) {
	// add a new T value to the collection
	valuesLength := len(tc.Values) // get the current length of the Values slice
	if valuesLength == 0 {
		// if the Values slice is empty, create a new slice of T with capacity X and append it to the Values slice
		tc.Values = append(tc.Values, make([]T, 0, tc.perArrayAmount))
		valuesLength = 1
	}
	lastSlice := len(tc.Values[valuesLength-1]) // get the current length of the last slice of T in the Values slice
	if lastSlice == tc.perArrayAmount {
		// if the last slice of T is full, create a new slice of T with capacity X and append it to the Values slice
		tc.Values = append(tc.Values, make([]T, 0, tc.perArrayAmount))
		valuesLength++ // update the length of the Values slice
	}
	// append the new T value to the last slice of T in the Values slice
	tc.Values[valuesLength-1] = append(tc.Values[valuesLength-1], t)
}

func AllValuesUnique[T comparable](array *[]T) bool {
	for i := 0; i < len(*array); i++ {
		if Index(array, (*array)[i], i+1) != -1 {
			return false
		}
	}
	return true
}

// Index returns the index of the first occurrence of v in s, after i,
// or -1 if not present.
func Index[S ~[]E, E comparable](s *S, v E, i int) int {
	for ; i < len(*s); i++ {
		if v == (*s)[i] {
			return i
		}
	}
	return -1
}
