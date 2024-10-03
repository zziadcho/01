package functions

import "math"

func StdDeviation(data []string) float64 {
	variance := Variance(data)
	return math.Sqrt(variance)
}
