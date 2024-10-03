package functions

import (
	"sort"
	"strconv"
)

func Median(data []string) float64 {
	if len(data) == 0 {
		return 0
	}
	var intList []float64
	for _, c := range data {
		if val, err := strconv.ParseFloat(string(c), 64); err == nil {
			intList = append(intList, val)
		}
	}

	sort.Float64s(intList)
	i := len(intList) / 2
	if i%2 == 1 {
		return intList[i]
	}
	return (intList[i] + intList[i-1]) / 2
}
