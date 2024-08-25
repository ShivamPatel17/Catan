package internal

import "math"

func HeightOfEqualateralTriangle(sideLength float64) float64 {
	return math.Sqrt(3) * sideLength / 2.0
}
