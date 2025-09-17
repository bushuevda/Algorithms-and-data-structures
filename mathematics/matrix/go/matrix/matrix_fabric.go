package matrix

import "errors"

type MatrixFabric[T Number] struct{}

func (mf *MatrixFabric[T]) Create(rowsSize int, colSize int) *Matrix[T] {

	arr := make([][]T, rowsSize)

	for i := range rowsSize {
		arr[i] = make([]T, colSize)
	}

	return &Matrix[T]{RowsCount: rowsSize, ColumnsCount: colSize, Arr2D: arr}
}

func (mf *MatrixFabric[T]) CreateByArr2D(arr2D [][]T) (*Matrix[T], error) {
	if len(arr2D) == 0 {
		return nil, errors.New("arr2D empty")

	}
	newMatrix := &Matrix[T]{RowsCount: len(arr2D), ColumnsCount: len(arr2D[0]), Arr2D: arr2D}

	return newMatrix, nil

}
