package line

import "math"

func DDA_line(x1 float64, y1 float64, x2 float64, y2 float64) ([]float64, []float64) {
	var i int = 0
	var x []float64
	var y []float64
	var x_start = math.Round(x1)
	var y_start = math.Round(y1)
	var x_end = math.Round(x2)
	var y_end = math.Round(y2)
	var length = max(math.Abs(x_end-x_start), math.Abs((y_end - y_start)))

	var dX = (x2 - x1) / length
	var dY = (y2 - y1) / length

	x = append(x, x1)
	y = append(y, y1)
	i++

	for i < int(length) {
		x = append(x, x[i-1]+dX)
		y = append(y, y[i-1]+dY)
		i++
	}
	x = append(x, x2)
	y = append(y, y2)

	return x, y
}
