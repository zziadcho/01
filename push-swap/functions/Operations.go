package functions

import "fmt"

/************************* Push Method ***************************/
func (s *Stack) Push(target *Stack) {
	if len(*s) == 0 {
		fmt.Println("error: source stack is empty")
		return
	}
	val := (*s)[0]
	*s = (*s)[1:]
	*target = append([]int{val}, *target...)
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