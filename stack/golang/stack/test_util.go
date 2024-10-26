package stack

/*
Data for testing all stack functions
*/
var TEST_DATA1 = []int{1, 2, 3, 4, 5, 6, 7}
var TEST_DATA2 = []string{"a", "e", "b", "d", "f"}
var TEST_DATA3 = []int{7, 5, 1, 4, 15, 6, 7}

/*
Data for testing stack function pop
*/
var TEST_DATA_POP1 = []int{1, 2, 3, 4}        // -3
var TEST_DATA_POP2 = []string{"a", "e", "b"}  // -2
var TEST_DATA_POP3 = []int{7, 5, 1, 4, 15, 6} //- 1

type testStackUtil[T comparable] struct{}

/*
Helped function for testing the push stack function
*/
func (t testStackUtil[T]) testPush(array []T) []T {
	var stack Stack[T]
	for _, element := range array {
		stack.Push(element)
	}
	return stack.array
}

/*
Helped function for testing the pop stack function
*/
func (t testStackUtil[T]) testPop(array []T, size_pop int) []T {
	var stack Stack[T]
	for _, element := range array {
		stack.Push(element)
	}

	for i := 0; i < size_pop; i++ {
		stack.Pop()
	}
	return stack.array
}

/*
Helped function for testing the top stack function
*/
func (t testStackUtil[T]) testTop(array []T) T {
	var stack Stack[T]
	for _, element := range array {
		stack.Push(element)
	}
	temp, _ := stack.Top()
	return temp
}

/*
Helped function for testing the size stack function
*/
func (t testStackUtil[T]) testSize(array []T) int {
	var stack Stack[T]

	for _, element := range array {
		stack.Push(element)
	}
	return stack.Size()
}

/*
Helped function for testing the empty stack function
*/
func (t testStackUtil[T]) testEmpty(array []T) bool {
	var stack Stack[T]

	for _, element := range array {
		stack.Push(element)
	}
	return stack.Empty()
}

/*
Helped function for testing the swap stack function
*/
func (t testStackUtil[T]) testSwap(array1 []T, array2 []T) []T {
	var stack1 Stack[T]
	for _, element := range array1 {
		stack1.Push(element)
	}

	var stack2 Stack[T]
	for _, element := range array2 {
		stack2.Push(element)
	}

	temp := stack1.array
	stack1.array = stack2.array
	stack2.array = temp
	return stack1.array
}

/*
Helpder function for comparate arrays
*/
func (t testStackUtil[T]) assertEqualsArrays(array1 []T, array2 []T) bool {
	var size int
	if len(array1) < len(array2) {
		size = len(array1)
	} else {
		size = len(array2)
	}

	for i := 0; i < size; i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

/*
Helpder function for comparate values
*/
func (t testStackUtil[T]) assertEqual(value1 T, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}
