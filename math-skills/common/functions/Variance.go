package functions

import (
	"math"
	"strconv"
)

func Variance(data []string) float64 {
	avg := Average(data)

	var intList []float64
	for _, c := range data {
		if val, err := strconv.ParseFloat(string(c), 64); err == nil {
			intList = append(intList, val)
		}
	}
	var squaredSum float64
	for _, num := range intList {
		diff := float64(num) - float64(avg)
		squaredSum += diff * diff
	}

	return math.Round(squaredSum / float64(len(intList))) 
}
