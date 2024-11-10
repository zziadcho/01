package functions

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

/************************* FindMaxIndex Method ***************************/
func (s *Stack) FindMaxIndex() int {
	maxIndex := 0
	for i := range *s {
		if (*s)[i] > (*s)[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

/************************* CalculateCost Function ***************************/
func CalculateCost(stack Stack, index int) int {
	length := len(stack)
	if index <= length/2 {
		return index
	}
	return length - index

}

/************************* IsSorted Function ***************************/
func IsSorted(stack []int) bool {
    for i := 1; i < len(stack); i++ {
        if stack[i-1] > stack[i] {
            return false
        }
    }
    return true
}

/************************* HasDuplicates ***************************/
func HasDuplicates(stack Stack) bool {
	seen := make(map[int]bool)
	for _, value := range stack {
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	return false
}
