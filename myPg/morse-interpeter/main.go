package main

import (
	"fmt"
	"01/morse-interpeter/common/functions"
)

func main() {
	var userInput string
	fmt.Println("enter a text to start translating")
	fmt.Scanln(&userInput)
	fmt.Printf("translation complete: %v \n",functions.MorseCodeInterpeter(userInput))
}