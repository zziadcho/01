package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

var TotalOps int

/************************* Push Method ***************************/
func (s *Stack) Push(target *Stack) {
	if len(*s) == 0 {
		fmt.Println("error: source stack is empty")
		return
	}
	val := (*s)[0]
	*s = (*s)[1:]
	*target = append(*target, val)
	TotalOps++
}

/************************* Swap Method ***************************/
func (s *Stack) Swap() {
	if len(*s) < 2 {
		return
	}
	(*s)[0], (*s)[1] = (*s)[1], (*s)[0]
	TotalOps++
	fmt.Println("sa") 
}

/************************* SwapBoth Method ***************************/
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
	TotalOps++
	fmt.Println("ra") 
}

/************************* RotateAll Method ***************************/
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
	TotalOps++
	fmt.Println("rra")
}

/************************* ReverseRotateAll Method ***************************/
func ReverseRotateAll(a, b *Stack) {
	a.ReverseRotate()
	b.ReverseRotate()
	fmt.Println("rrr")
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
func (s *Stack) CalculateCost(index int) int {
	length := len(*s)
	if index <= length/2 {
		return index
	}
	return length - index
}

/************************* PerformCheapestOperation Method ***************************/
func (s *Stack) PerformCheapestOperation() {
	minIndex := s.FindMinIndex()
	cost := s.CalculateCost(minIndex)

	if minIndex <= len(*s)/2 {
		for i := 0; i < cost; i++ {
			s.Rotate()
		}
	} else {
		for i := 0; i < cost; i++ {
			s.ReverseRotate()
		}
	}
}

/************************* PushBackToStackA Method ***************************/
func PushBackToStackA(stackA, stackB *Stack) {
	for len(*stackB) > 0 {
		stackB.PerformCheapestOperation()
		stackB.Push(stackA)
		fmt.Println("pa") 
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
	fmt.Println("Before sorting: ", stackA)

	for len(stackA) > 0 {
		stackA.PerformCheapestOperation()
		stackA.Push(&stackB)
		fmt.Println("pb")
	}
	
	fmt.Println("After sorting: ", stackB)
	fmt.Printf("Total Operations: %v\n", TotalOps)
}
