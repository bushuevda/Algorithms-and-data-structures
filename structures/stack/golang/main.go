package main

import (
	"stack/stack"
)

func main() {
	b := stack.Stack[int16]{}
	b.Push(1)
	b.Push(2)
	b.Push(3)

	a := stack.Stack[int16]{}
	a.Push(4)
	a.Push(5)
	a.Push(6)
	a.Push(6)
	a.Swap(&b)
	a.Show()
}
