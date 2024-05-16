package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	standard, _ := os.ReadFile("standard.txt")
	standardString := string(standard)
	standardSplit := strings.Split(standardString, "\n\n")
	letter :=  standardSplit[65]
	letterSplit := strings.Split(letter,"\n")
	fmt.Println(letterSplit)

}
