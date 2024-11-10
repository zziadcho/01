package main

import (
	"fmt"
	"os"
	"push-swap/functions"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
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

	if functions.HasDuplicates(stackA) {
		fmt.Println("Error")
		return
	}

	if functions.IsSorted(stackA) {
		fmt.Println("Already Sorted")
		return
	}

	stackB := functions.Stack{}

	fmt.Println("Initial stacks:", stackA, stackB)

	switch {
	case len(stackA) == 2:
		functions.IbogaSort2Numbers(&stackA)
	case len(stackA) == 3:
		functions.IbogaSort3Numbers(&stackA)
	case len(stackA) >= 5:
		functions.IbogaSortChunks(&stackA, &stackB)
	}

	functions.IbogaSortChunks(&stackA, &stackB)

	fmt.Println("Final stacks:", stackA, stackB)
	fmt.Println("Total Operations:", functions.TotalOperations)
}
