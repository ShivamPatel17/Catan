package internal

import "math"

func HeightOfEqualateralTriangle(sideLength float64) float64 {
	return math.Sqrt(3) * sideLength / 2.0
}

func WithinTolerance(x float64, y float64, t float64) bool {
	return math.Abs(y-x) < t
}
