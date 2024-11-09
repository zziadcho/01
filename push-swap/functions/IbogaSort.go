package functions

import "fmt"

type Stack []int

var TotalOperations int

/************************* IbogaSort3Numbers ***************************/
func IbogaSort2Numbers(stack *Stack){
	if (*stack)[0] > (*stack)[1]{
		
		stack.Swap()
	} else {
	}
}
/************************* IbogaSort3Numbers ***************************/

func IbogaSort3Numbers(stackA *Stack) {

	if len(*stackA) == 0 {
		fmt.Println("Error")
	}
	top, mid, bot := (*stackA)[0], (*stackA)[1], (*stackA)[2]

	if bot > top && top > mid {
		(*stackA).Swap()
		fmt.Println("sa")

	} else if top > mid && mid > bot {
		(*stackA).Swap()
		(*stackA).ReverseRotate()
		fmt.Println("sa")
		fmt.Println("rra")

	} else if top > bot && bot > mid {
		(*stackA).Rotate()
		fmt.Println("ra")

	} else if mid > bot && bot > top {
		(*stackA).Swap()
		(*stackA).Rotate()
		fmt.Println("sa")
		fmt.Println("ra")

	} else if bot < top && top < mid {
		(*stackA).ReverseRotate()
		fmt.Println("rra")
	}
}

/************************* IbogaSort5Numbers ***************************/
func IbogaSort5Numbers(stackA, stackB *Stack) {
	if len((*stackA)) == 0 {
		fmt.Println("Error")
	}
	
	(*stackA).Push(stackB)
	(*stackA).Push(stackB)

	IbogaSort3Numbers(stackA)
	IbogaSort2Numbers(stackB)
	for len(*stackB) > 0 {
		(*stackB).Push(stackA)
	}
}

/************************* IbogaSort100Numbers ***************************/
func IbogaSort100Numbers(stackA, stackB *Stack) {
	if len((*stackA)) == 0 {
		fmt.Println("Error")
	}

	chunkLength := len(*stackA) / 5
	var firstHoldCost, secondHoldCost, chunkNumber int

	for len((*stackA)) > 0 {
		if len(*stackB) != 0 {
			chunkNumber = len(*stackB) / chunkLength
		}
		for i := 0; i < len((*stackA)); i++ {
			if (*stackA)[i] >= chunkLength*chunkNumber && (*stackA)[i] <= chunkLength*(chunkNumber+1) {
				firstHoldCost = i
				break
			}
		}

		for i := len((*stackA)) - 1; i > 0; i-- {
			if (*stackA)[i] >= chunkLength*chunkNumber && (*stackA)[i] <= chunkLength*(chunkNumber+1) {
				secondHoldCost = len((*stackA)) - i
				break
			}
		}

		if firstHoldCost < secondHoldCost {
			for i := 0; i < firstHoldCost; i++ {
				(*stackA).Rotate()
				fmt.Println("ra")
				TotalOperations++
			}
		} else {
			for i := 0; i < secondHoldCost; i++ {
				(*stackA).ReverseRotate()
				fmt.Println("rra")
				TotalOperations++
			}
		}

		if len(*stackB) == 0 {
			(*stackA).Push(stackB)
			fmt.Print("pb")
			TotalOperations++
			continue
		}

		minIndex := stackB.FindMinIndex()
		costMinNum := CalculateCost(*stackB, minIndex)

		if costMinNum < len(*stackB) {
			for i := 0; i < costMinNum; i++ {
				stackB.Rotate()
				fmt.Println("rb")
				TotalOperations++
			}
		} else {
			for i := 0; i < costMinNum; i++ {
				stackB.ReverseRotate()
				fmt.Println("rrb")
				TotalOperations++
			}
		}

		(*stackA).Push(stackB)
		fmt.Println("pb")
		TotalOperations++
	}

	for len(*stackB) > 0 {
		maxIndex := stackB.FindMaxIndex()
		costMaxNum := CalculateCost(*stackB, maxIndex)

		if maxIndex > len(*stackB)/2 {
			for i := 0; i < costMaxNum; i++ {
				stackB.ReverseRotate()
				fmt.Println("rrb")
				TotalOperations++
			}
		} else {
			for i := 0; i < costMaxNum; i++ {
				stackB.Rotate()
				fmt.Println("rb")
				TotalOperations++
			}
		}
		(*stackB).Push(stackA)
		fmt.Println("pa")
		TotalOperations++
	}
}
