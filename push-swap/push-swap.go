package main

import (
	"fmt"
	"os"
	"push-swap/functions"
	"strconv"
	"strings"
)

/************************* Main ***************************/
func main() {
	if len(os.Args) < 2 {
		return
	}

	input := strings.Split(os.Args[1], " ")

	stackA := functions.Stack{}
	for _, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error")
			return
		}
		stackA = append(stackA, num)
	}
	stackB := functions.Stack{}

	if functions.IsSorted(stackA) {
		fmt.Println("Already Sorted")
		return
	}
	fmt.Println(stackA, stackB)
	if len(stackA) == 2 {
		functions.IbogaSort2Numbers(&stackA)
	} else if len(stackA) == 3 {
		functions.IbogaSort3Numbers(&stackA)
	} else if len(stackA) <= 5 {
		functions.IbogaSort5Numbers(&stackA, &stackB)
	}
	
	fmt.Println(stackA, stackB)
	fmt.Println(functions.TotalOperations)
}
