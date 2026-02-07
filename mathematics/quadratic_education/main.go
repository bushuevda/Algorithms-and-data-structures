package main

import (
	mathEquation "QuadraticEcuation/math_equations/quadratic_equation"
	"fmt"
)

func main() {
	qe := mathEquation.QuadraticEquation{}
	x1, x2 := qe.Calculate(-1, 6, 7)
	fmt.Println(x1, x2)
}
