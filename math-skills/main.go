package main

import (
	"01/math-skills/common/functions"
	"fmt"
	"math"
)

func main() {
	functions.StartUp()
	fmt.Println("Average:", int(math.Round(functions.Average(functions.Numbers))))
	fmt.Println("Median:", int(math.Round(functions.Median(functions.Numbers))))
	fmt.Println("Variance:", int(math.Round(functions.Variance(functions.Numbers))))
	fmt.Println("Standard Deviation:", int(math.Round(functions.StdDeviation(functions.Numbers))))
}
