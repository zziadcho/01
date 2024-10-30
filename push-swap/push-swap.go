package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

/************************* Push Method ***************************/
func (s *Stack) Push(target *Stack) {
	if len(*s) == 0 {
		fmt.Println("error: source stack is empty")
		return
	}
	val := (*s)[0]
	*s = (*s)[1:]
	*target = append(*target, val)
}

/************************* Swap Method ***************************/
func (s *Stack) Swap() {
	if len(*s) < 2 {
		return
	}
	(*s)[0], (*s)[1] = (*s)[1], (*s)[0]
}

/************************* SwapBoth Function ***************************/
func SwapAll(a, b *Stack) {
	a.Swap()
	b.Swap()
	fmt.Println("ss")
}

/************************* Rotate Method ***************************/
func (s *Stack) Rotate() {
	if len(*s) < 2 {
		return
	}
	top := (*s)[0]
	for i := 1; i < len(*s); i++ {
		(*s)[i-1] = (*s)[i]
	}
	(*s)[len(*s)-1] = top
}

/************************* RotateAll Function ***************************/
func RotateAll(a, b *Stack) {
	a.Rotate()
	b.Rotate()
	fmt.Println("rr")
}

/************************* ReverseRotate Method ***************************/
func (s *Stack) ReverseRotate() {
	if len(*s) < 2 {
		return
	}
	bottom := (*s)[len(*s)-1]
	for i := len(*s) - 1; i > 0; i-- {
		(*s)[i] = (*s)[i-1]
	}
	(*s)[0] = bottom
}

/************************* ReverseRotateAll Function ***************************/
func ReverseRotateAll(a, b *Stack) {
	a.ReverseRotate()
	b.ReverseRotate()
}

/************************* FindMinIndex Method ***************************/
func (s *Stack) FindMinIndex() int {
	minIndex := 0
	for i := range *s {
		if (*s)[i] < (*s)[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

/************************* CalculateCost Method ***************************/
func CalculateCost(stack Stack, index int) int {
	length := len(stack)
	if index <= length/2 {
		return index
	}
	return length - index

}

/************************* SortByChunk Function ***************************/
func IbogaSort100Numbers(stackA, stackB *Stack) {
	if len((*stackA)) == 0 {
		os.Exit(1)
	}
	var firstHoldCost, secondHoldCost int // len(stackA) / 5 mn moraha ndir (len(stackA)/5 * 2) dik zoj ghadi tb9a tzid
	for len((*stackA)) > 0 {
		for i := 0; i < len((*stackA)); i++ {
			if (*stackA)[i] >= 0 && (*stackA)[i] <= 19 {
				firstHoldCost = i
				break
			}
		}
	
		for i := len((*stackA))-1; i > 0; i-- {
			if (*stackA)[i] >= 0 && (*stackA)[i] <= 19 {
				secondHoldCost = len((*stackA))-i
				break
			}
		}
	
		if firstHoldCost < secondHoldCost {
			for i := 0; i < firstHoldCost; i++ {
				(*stackA).Rotate()
				fmt.Println("ra")
			}
		} else {
			for i := 0; i < secondHoldCost; i++ {
				(*stackA).ReverseRotate()
				fmt.Println("rra")
			}
		}
	
		if len(*stackB) == 0 {
			(*stackA).Push(stackB)
			fmt.Print("pb")
			continue
		}
	
		minIndex := stackB.FindMinIndex()
		costNewNum := CalculateCost(*stackB, minIndex)
	
		if costNewNum < len(*stackB) {
			for i := 0; i < costNewNum; i++ {
				stackB.Rotate()
				fmt.Println("rb")
			}
		} else {
			for i := 0; i < costNewNum; i++ {
				stackB.ReverseRotate()
				fmt.Println("rb")
			}
		}

			(*stackA).Push(stackB)
			fmt.Println("pb")
	}
}

/************************* Main ***************************/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: no input")
		os.Exit(1)
	}

	input := strings.Split(os.Args[1], " ")

	stackA := Stack{}
	for _, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error: converting to integer")
			os.Exit(1)
		}
		stackA = append(stackA, num)
	}

	stackB := Stack{}

	println("----------")
	fmt.Println("before sorting:" , stackA, stackB)
	println("----------")

	IbogaSort100Numbers(&stackA, &stackB)

	println("----------")
	fmt.Println("after sorting:" , stackA, stackB)
	println("----------")
}
