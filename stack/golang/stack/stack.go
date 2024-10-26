package stack

import (
	"fmt"
)

type Stack[T any] struct {
	array []T
}

/*
A function that adds an element to the stack
*/
func (s *Stack[T]) Push(element T) {
	s.array = append(s.array, element)
}

/*
A function that removes a top element from the stack and returns it
*/
func (s *Stack[T]) Pop() (T, error) {

	if s.Size() > 0 {
		e := s.array[len(s.array)-1]
		s.array = s.array[:len(s.array)-1]
		return e, nil
	} else {
		var b T
		return b, fmt.Errorf("error of the pop function, the stack size must be more than %d ", s.Size())
	}
}

/*
A function that returns stack size
*/
func (s *Stack[T]) Size() int {
	return len(s.array)
}

/*
A function that checks the stack for emptiness
*/
func (s *Stack[T]) Empty() bool {
	if s.Size() > 0 {
		return true
	} else {
		return false
	}
}

/*
A function that returns the top element of the stack
*/
func (s *Stack[T]) Top() (T, error) {
	if s.Size() > 0 {
		return s.array[len(s.array)-1], nil
	} else {
		var b T
		return b, fmt.Errorf("error of the top function, the stack size must be more than %d ", s.Size())
	}
}

/*
A function that swaps the stack arrays
*/
func (s *Stack[T]) Swap(s2 *Stack[T]) {
	temp := s2.array
	s2.array = s.array
	s.array = temp
}

/*
A function that shows the stack arrays
*/
func (s *Stack[T]) Show() {
	for _, element := range s.array {
		fmt.Print(element, " ")
	}
}
