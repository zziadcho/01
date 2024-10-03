package functions

import (
	"math"
	"strconv"
)

func Average(data []string) float64 {
	if len(data) == 0 {
		return 0
	}

	var sum int
	var count int

	for _, c := range data {
		if val, err := strconv.Atoi(string(c)); err == nil {
			sum += val
			count++
		}
	}
	if count == 0 {
		return 1
	}

	return math.Round(float64(sum) / float64(count))
}
