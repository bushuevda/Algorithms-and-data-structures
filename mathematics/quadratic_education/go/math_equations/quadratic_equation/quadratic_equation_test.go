package math_equations

import (
	"math"
	"testing"
)

const EPS = 1.0e-4

func compare_float(a float32, b float32, eps float32) bool {
	return float32(math.Abs(float64(a-b))) < eps
}

func util_test_calc_abc(a float32, b float32, c float32, x1 float32, x2 float32) bool {
	qe := QuadraticEquation{}

	var D float32 = b*b - 4*a*c
	x1_, x2_ := qe.Calculate(a, b, c)

	if D > 0 && compare_float(x1_, x1, EPS) && compare_float(x2_, x2, EPS) {
		return true
	} else if D == 0 && compare_float(x1_, x1, EPS) && math.IsNaN(float64(x2_)) && math.IsNaN(float64(x2)) {
		return true
	} else if D < 0 && math.IsNaN(float64(x1)) && math.IsNaN(float64(x1_)) && math.IsNaN(float64(x2)) && math.IsNaN(float64(x2_)) {
		return true
	}

	return false
}

func TestCalculateTwoRoots(t *testing.T) {
	var test_data = [][]float32{
		//a = -1, b = 6, c = 7, x1 = 7, x2 = -1
		{-1, 6, 7, 7, -1},
		//a = -1, b = 5, c = -6, x1 = -6, x2 = 1
		{1, 5, -6, -6, 1},
		//a = 1, b = 1, c = -6, x1 = -3, x2 = 2
		{1, 1, -6, -3, 2},
		//a = 4, b = 1, c = -1, x1 = -0.64039, x2 = 0.39039
		{4, 1, -1, -0.64039, 0.39039},
		//a = 4, b = 11, c = 3, x1 = -2.4430, x2 = -0.30700
		{4, 11, 3, -2.4430, -0.30700},
	}
	for i, arr := range test_data {
		if !util_test_calc_abc(arr[0], arr[1], arr[2], arr[3], arr[4]) {
			t.Error("\nTest error for data with index ", i)
		}
	}

}

func TestCalculateOneRoot(t *testing.T) {
	var test_data = [][]float32{
		//a = 9, b = -6, c = 1, x1 =  0.33333, x2 = NaN
		{9, -6, 1, 0.33333, float32(math.NaN())},
		//a = 4, b = 4, c = 1, x1 = -0.5, x2 = NaN
		{4, 4, 1, -0.5, float32(math.NaN())},
		//a = 1, b = -6, c = 9, x1 = 3, x2 = NaN
		{1, -6, 9, 3, float32(math.NaN())},
		//a = 1, b = 12, c = 36, x1 = -6, x2 = NaN
		{1, 12, 36, -6, float32(math.NaN())},
		//a = 1, b = 4, c = 4, x1 = -2, x2 = NaN
		{1, 4, 4, -2, float32(math.NaN())},
	}
	for i, arr := range test_data {
		if !util_test_calc_abc(arr[0], arr[1], arr[2], arr[3], arr[4]) {
			t.Error("\nTest error for data with index ", i)
		}
	}

}

func TestWithOutRoots(t *testing.T) {
	var test_data = [][]float32{
		//a = 4, b = 2, c = 3, x1 = NaN, x2 = NaN
		{4, 2, 3, float32(math.NaN()), float32(math.NaN())},
		//a = 7, b = 6, c = 3, x1 = NaN, x2 = NaN
		{7, 6, 3, float32(math.NaN()), float32(math.NaN())},
		//a = 7, b = 6, c = 7, x1 = NaN, x2 = NaN
		{7, 6, 7, float32(math.NaN()), float32(math.NaN())},
		//a = -11, b = 1, c = -3, x1 = NaN, x2 = NaN
		{-11, 1, -3, float32(math.NaN()), float32(math.NaN())},
		//a = 21, b = -11, c = 5, x1 = NaN, x2 = NaN
		{21, -11, 5, float32(math.NaN()), float32(math.NaN())},
	}
	for i, arr := range test_data {
		if !util_test_calc_abc(arr[0], arr[1], arr[2], arr[3], arr[4]) {
			t.Error("\nTest error for data with index ", i)
		}
	}

}

func TestA(t *testing.T) {
	var test_data = [][]float32{
		//a = -1, b = 0, c = 0, x1 = 0, x2 = NaN
		{-1, 0, 0, 0, float32(math.NaN())},
		//a = 5, b = 0, c = 0, x1 = 0, x2 = NaN
		{5, 0, 0, 0, float32(math.NaN())},
		//a = 12, b = 0, c = 0, x1 = 0, x2 = NaN
		{12, 0, 0, 0, float32(math.NaN())},
		//a = 4, b = 0, c = 0, x1 = 0, x2 = NaN
		{4, 0, 0, 0, float32(math.NaN())},
		//a = 7, b = 0, c = 0, x1 = 0, x2 = NaN
		{7, 0, 0, 0, float32(math.NaN())},
	}
	for i, arr := range test_data {
		if !util_test_calc_abc(arr[0], arr[1], arr[2], arr[3], arr[4]) {
			t.Error("\nTest error for data with index ", i)
		}
	}

}

func TestAB(t *testing.T) {
	var test_data = [][]float32{
		//a = 12, b = 4, c = 0, x1 = -0.33333, x2 = 0
		{12, 4, 0, -0.33333, 0},
		//a = 5, b = -4, c = 0, x1 = 0, x2 = 0.8
		{5, -4, 0, 0, 0.8},
		//a = 5, b = 5, c = 0, x1 = -1, x2 = 0
		{5, 5, 0, -1, 0},
		//a = 13, b = 15, c = 0, x1 = -1.1538, x2 = 0
		{13, 15, 0, -1.1538, 0},
		//a = -13, b = 15, c = 0, x1 = 1.1538, x2 = 0
		{-13, 15, 0, 1.1538, 0},
	}
	for i, arr := range test_data {
		if !util_test_calc_abc(arr[0], arr[1], arr[2], arr[3], arr[4]) {
			t.Error("\nTest error for data with index ", i)
		}
	}

}

func TestAC(t *testing.T) {
	var test_data = [][]float32{
		//a = -1, b = 0, c = 7, x1 = 2.6458, x2 = -2.6458
		{-1, 0, 7, 2.6458, -2.6458},
		//a = 7, b = 0, c = 7, x1 = NaN, x2 = NaN
		{7, 0, 7, float32(math.NaN()), float32(math.NaN())},
		//a = 8, b = 0, c = 7, x1 = NaN, x2 = NaN
		{8, 0, 7, float32(math.NaN()), float32(math.NaN())},
		//a = 1, b = 0, c = -17, x1 = -4.1231, x2 = 4.1231
		{1, 0, -17, -4.1231, 4.1231},
		//a = 3, b = 0, c = -5, x1 = -1.2910, x2 = 1.2910
		{3, 0, -5, -1.2910, 1.2910},
	}
	for i, arr := range test_data {
		if !util_test_calc_abc(arr[0], arr[1], arr[2], arr[3], arr[4]) {
			t.Error("\nTest error for data with index ", i)
		}
	}

}
