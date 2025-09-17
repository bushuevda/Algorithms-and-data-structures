package main

import (
	"fmt"
	"matrix/matrix"
)

func TM(arr [][]int32) {

}

func main() {
	mf := matrix.MatrixFabric[int32]{}

	m1, err := mf.CreateByArr2D([][]int32{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	})

	if err != nil {
		fmt.Errorf("%s", err)
		return
	}

	m2, err := mf.CreateByArr2D([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	if err != nil {
		fmt.Errorf("%s", err)
		return
	}

	m1.Add(m2)

	m3, err := m2.Transporate()

	if err != nil {
		fmt.Errorf("%s", err)
		return
	}

	m3.ShowMatrix()

}
