package math_equations

import "math"

type QuadraticEquation struct{}

func (qe QuadraticEquation) Calculate(a float32, b float32, c float32) (float32, float32) {
	var D float32 = b*b - 4*a*c
	var x1, x2 float32

	if D > 0 {
		x1 = (-b - float32(math.Sqrt(float64(D)))) / (2 * a)
		x2 = (-b + float32(math.Sqrt(float64(D)))) / (2 * a)
	} else if D == 0 {
		x1 = -b / (2 * a)
		x2 = float32(math.NaN())
	} else if D < 0 {
		x1 = float32(math.NaN())
		x2 = float32(math.NaN())
	}

	return x1, x2
}
