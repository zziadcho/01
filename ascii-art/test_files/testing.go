package main

import (
	"fmt"
	"os"
	"regexp"	
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run your_program.go <string>")
        return
    }

	target := regexp.MustCompile(`\\n`)
    printables := os.Args[1]
	lockedIn := target.FindStringIndex(printables)
	beforeNewLine := printables[:lockedIn[0]]
	newLine := printables[lockedIn[0]:lockedIn[1]]
	afterNewLine := printables[lockedIn[1]:]
	fmt.Println(beforeNewLine,newLine, afterNewLine)
}
