package main

import (
	"fmt"
	"graphics_algorithms/line"
)

func main() {
	x, y := line.DDA_line(1, 2, 3, 4)

	for i := 0; i < len(x); i++ {
		fmt.Print(x[i])
	}
	fmt.Println()
	for i := 0; i < len(y); i++ {
		fmt.Print(y[i])
	}
	fmt.Println()
}
