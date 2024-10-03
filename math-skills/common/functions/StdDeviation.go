package functions

import "math"

func StdDeviation(data []string) interface{} {
	variance := Variance(data)
	return math.Round(math.Sqrt(variance)) 
}
