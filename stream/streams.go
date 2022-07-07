package stream

import (
	"math/rand"
	"reflect"
)

var (
	random = rand.New(rand.NewSource(rand.Int63()))
)

type Streamable[T any] interface {
	RemoveIndex(int) Streamable[T]
	Filter(func(T)) Streamable[T]
	Contains(T) bool
	ForEach(func(T))
	ForEachIndexed(func(int, T))
	Count(func(T) bool) uint64
	All(func(T) bool) bool
	Any(func(T) bool) bool
	NotEmpty() bool
	Size() int
	Find(func(T) bool) T
	Last(func(T) bool) T
	Random() T
	Drop(int) []T
	DropLast(int) ([]T, error)
	DropLastWhile(func(T) bool) ([]T, error)
	Slice([]int) []T
	DropWhile(func(T) bool) []T
	Take() (Streamable[T], error)
	TakeLast(int) (Streamable[T], error)
	TakeWhile(func(T) bool) Streamable[T]
	Map() Streamable[any]
}

type Stream[T any] struct {
	elements []T
}

func Of[T any](arr ...T) *Stream[T] {
	return &Stream[T]{elements: arr}
}

// RemoveIndex removes the provided index from arr and returns an updated array.
func (s *Stream[T]) RemoveIndex(index int) *Stream[T] {
	s.elements = append(s.elements[:index], s.elements[index+1:]...)
	return s
}

// Filter filters the output based on the provided predicate.
func (s *Stream[T]) Filter(predicate func(elem T) bool) *Stream[T] {
	var result []T
	for _, elem := range s.elements {
		if predicate(elem) {
			result = append(result, elem)
		}
	}
	s.elements = result
	return s
}

// Contains returns true if element is found in the arr.
func (s *Stream[T]) Contains(element T) bool {
	for _, elem := range s.elements {
		if reflect.DeepEqual(elem, element) {
			return true
		}
	}
	return false
}

// ForEach performs the given action on each element.
func (s *Stream[T]) ForEach(function func(T)) {
	for _, elem := range s.elements {
		function(elem)
	}
}

// ForEachIndexed performs the given action on each element, providing sequential index with the element.
func (s *Stream[T]) ForEachIndexed(action func(index int, element T)) {
	for i, elem := range s.elements {
		action(i, elem)
	}
}

// Count returns the count of elements matching the given predicate.
func (s *Stream[T]) Count(predicate func(elem T) bool) (count uint64) {
	for _, elem := range s.elements {
		if predicate(elem) {
			count++
		}
	}
	return count
}

// All returns true if all elements match the given predicate.
func (s *Stream[T]) All(predicate func(elem T) bool) bool {
	for _, elem := range s.elements {
		if !predicate(elem) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element matches the given predicate.
func (s *Stream[T]) Any(predicate func(elem T) bool) bool {
	for _, elem := range s.elements {
		if predicate(elem) {
			return true
		}
	}
	return false
}

// NotEmpty returns true if arr has at least one element.
func (s *Stream[T]) NotEmpty() bool {
	return s.Size() > 0
}

// Empty returns true if arr has zero elements.
func (s *Stream[T]) Empty() bool {
	return s.Size() == 0
}

// Size returns the size of this Stream.
func (s *Stream[T]) Size() (count int) {
	return len(s.elements)
}

// Find finds the first element matching the given predicate, or nil if no such element was found.
func (s *Stream[T]) Find(predicate func(elem T) bool) *T {
	for _, elem := range s.elements {
		if predicate(elem) {
			return &elem
		}
	}
	return nil
}

// Last finds the last element matching the given predicate, or nil if no such element was found.
func (s *Stream[T]) Last(predicate func(elem T) bool) (element T) {
	for _, elem := range s.elements {
		if predicate(elem) {
			element = elem
		}
	}
	return element
}

// Random returns a random element from arr.
func (s *Stream[T]) Random() T {
	return s.elements[random.Intn(s.Size()-1)]
}

// Drop returns an array containing all elements except the first n elements.
func (s *Stream[T]) Drop(n int) *Stream[T] {
	if n >= s.Size() {
		return Of[T]()
	}
	s.elements = s.elements[n:]
	return s
}

// DropLast returns an arr containing all elements except last n elements.
func (s *Stream[T]) DropLast(n int) *Stream[T] {
	return s.Take(coerceAtLeast(s.Size()-n, 0))
}

// DropLastWhile returns an arr containing all elements except the last elements that satisfy the given predicate.
func (s *Stream[T]) DropLastWhile(predicate func(elem T) bool) *Stream[T] {
	if s.NotEmpty() {
		for i := s.Size(); i > 0; i-- {
			if !predicate(s.elements[i]) {
				return s.Take(i)
			}
		}
	}
	return s
}

// Slice returns an array containing elements at the specified indices.
func (s *Stream[T]) Slice(indices []int) *Stream[T] {
	arrLen := len(indices)
	if arrLen == 0 {
		s.elements = []T{}
		return s
	}
	var result []T
	for _, index := range indices {
		if index < len(s.elements) {
			result = append(result, s.elements[index])
		}
	}
	s.elements = result
	return s
}

// DropWhile returns an array containing all elements except the first elements that satisfy the given predicate.
func (s *Stream[T]) DropWhile(predicate func(elem T) bool) *Stream[T] {
	var yielding = false
	var result []T
	for _, elem := range s.elements {
		if yielding {
			result = append(result, elem)
		} else if !predicate(elem) {
			result = append(result, elem)
			yielding = true
		}
	}
	s.elements = result
	return s
}

// Take returns an array containing first n elements.
func (s *Stream[T]) Take(n int) *Stream[T] {
	if n < 0 {
		s.elements = []T{}
		return s
	}
	if n == 0 {
		s.elements = []T{}
		return s
	}
	var result []T

	arrLen := s.Size()
	if n >= arrLen {
		result = append(s.elements)
		s.elements = result
		return s
	}
	if n == 1 {
		result = append(result, s.elements[0])
		s.elements = result
		return s
	}

	count := 0
	for _, elem := range s.elements {
		result = append(result, elem)
		count++
		if count == n {
			break
		}
	}
	s.elements = result
	return s
}

// TakeLast returns an array containing last n elements.
func (s *Stream[T]) TakeLast(n int) *Stream[T] {
	if n < 0 {
		panic("n is less than 0")
	}
	if n == 0 {
		s.elements = []T{}
		return s
	}
	arrLen := s.Size()
	var result []T
	if n >= arrLen {
		return s
	}
	if n == 1 {
		result = append(result, s.elements[arrLen-1])
		s.elements = result
		return s
	}
	for index := arrLen - n; index < len(s.elements); index++ {
		result = append(result, s.elements[index])
	}
	s.elements = result
	return s
}

// TakeWhile returns an array containing the first elements satisfying the given predicate.
func (s *Stream[T]) TakeWhile(predicate func(elem T) bool) *Stream[T] {
	var result []T
	for _, elem := range s.elements {
		if !predicate(elem) {
			break
		}
		result = append(result, elem)
	}
	s.elements = result
	return s
}

//Map returns a stream transformed with the provided function.
//Limited to transforming to the same type.
func (s *Stream[T]) Map(transform func(elem T) T) *Stream[T] {
	var result []T
	for _, elem := range s.elements {
		result = append(result, transform(elem))
	}
	s.elements = result
	return s
}

func Map[T any, O any](arr []T, transform func(elem T) O) []O {
	var result []O
	for _, elem := range arr {
		result = append(result, transform(elem))
	}
	return result
}

func (s *Stream[T]) Values() []T {
	return s.elements
}
