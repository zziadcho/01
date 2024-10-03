package main

import (
	"01/math-skills/common/functions"
	"fmt"
)

func main() {
	functions.StartUp()
	fmt.Println("Average:", functions.Average(functions.Numbers))
	fmt.Println("Median:", functions.Median(functions.Numbers))
	fmt.Println("Variance:", functions.Variance(functions.Numbers))
	fmt.Println("Standard Deviation:", functions.StdDeviation(functions.Numbers))
}
