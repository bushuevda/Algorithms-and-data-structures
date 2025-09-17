package matrix

import (
	"errors"
	"fmt"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 |
		~uint | ~uint32 | ~uint16 | ~uint64
}

type Matrix[T Number] struct {
	RowsCount    int
	ColumnsCount int
	Arr2D        [][]T
}

func (m *Matrix[T]) ShowMatrix() {
	for i := 0; i < m.RowsCount; i++ {
		for j := 0; j < m.ColumnsCount; j++ {
			fmt.Print(m.Arr2D[i][j], "\t")
		}
		fmt.Println()
	}
}

func (m *Matrix[T]) Add(m2 *Matrix[T]) (*Matrix[T], error) {
	if m.ColumnsCount != m2.ColumnsCount || m.RowsCount != m2.RowsCount {
		return nil, errors.New("matrix size columns not equil size rows")
	}

	mf := MatrixFabric[T]{}
	sum_m := mf.Create(m.RowsCount, m.ColumnsCount)

	for i := 0; i < m.RowsCount; i++ {
		for j := 0; j < m.ColumnsCount; j++ {
			sum_m.Arr2D[i][j] = m.Arr2D[i][j] + m2.Arr2D[i][j]
		}
	}

	return sum_m, nil
}

func (m *Matrix[T]) Sub(m2 *Matrix[T]) (*Matrix[T], error) {
	if m.ColumnsCount != m2.ColumnsCount || m.RowsCount != m2.RowsCount {
		return nil, errors.New("matrix size columns not equil size rows")
	}

	mf := MatrixFabric[T]{}
	sub_m := mf.Create(m.RowsCount, m.ColumnsCount)

	for i := 0; i < m.RowsCount; i++ {
		for j := 0; j < m.ColumnsCount; j++ {
			sub_m.Arr2D[i][j] = m.Arr2D[i][j] - m2.Arr2D[i][j]
		}
	}

	return sub_m, nil
}

func (m *Matrix[T]) Dot(m2 *Matrix[T]) (*Matrix[T], error) {
	if m.ColumnsCount != m2.RowsCount || m.RowsCount != m2.RowsCount {
		return nil, errors.New("matrix size columns not equil size rows")
	}

	dot_matrix := make([][]T, 0)

	for i := 0; i < m.RowsCount; i++ {
		n_row := make([]T, 0)
		for j := 0; j < m.ColumnsCount; j++ {
			var sum T
			for k := 0; k < m.RowsCount; k++ {
				sum += m.Arr2D[i][k] * m2.Arr2D[k][j]
			}
			n_row = append(n_row, sum)
		}
		dot_matrix = append(dot_matrix, n_row)
	}

	mf := MatrixFabric[T]{}
	res, err := mf.CreateByArr2D(dot_matrix)
	return res, err
}

func (m *Matrix[T]) Mul(scalar T) *Matrix[T] {
	mf := MatrixFabric[T]{}
	sub_m := mf.Create(m.RowsCount, m.ColumnsCount)

	for i := 0; i < m.RowsCount; i++ {
		for j := 0; j < m.ColumnsCount; j++ {
			sub_m.Arr2D[i][j] = m.Arr2D[i][j] * scalar
		}
	}

	return sub_m
}

func (m *Matrix[T]) Transporate() (*Matrix[T], error) {
	transporateMatrix := make([][]T, 0)

	for i := 0; i < m.ColumnsCount; i++ {
		t_column := make([]T, 0)
		for j := 0; j < m.RowsCount; j++ {
			t_column = append(t_column, m.Arr2D[j][i])
		}
		transporateMatrix = append(transporateMatrix, t_column)
	}

	mf := MatrixFabric[T]{}
	res, err := mf.CreateByArr2D(transporateMatrix)
	return res, err
}
