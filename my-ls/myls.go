package main

import (
	"fmt"
	"os"
)

type Flags struct {
	Recursive   bool
	Reverse     bool
	Hidden      bool
	All         bool
	Time        bool
	Long_Format bool
}

func main() {
	flags := Flags{}
	// getting the arguments and splitting them, after that determine what flags to use
	if len(os.Args) == 1 {
		// Function to list the current directory
		fmt.Println("Listing current directory with default options...")
		// listDirectory(".", flags)
	} else {
		args := os.Args[1:]
		for _, argument := range args {
			if len(argument) > 1 && argument[0] == '-' {
				for _, char := range argument[1:] {
					if char == 'R' {
						flags.Recursive = true
					} else if char == 'r' {
						flags.Reverse = true
					} else if char == 'h' {
						flags.Hidden = true
					} else if char == 'a' {
						flags.All = true
					} else if char == 't' {
						flags.Time = true
					} else if char == 'l' {
						flags.Long_Format = true
					}
				}
			}

			fmt.Println(flags)

		}

		// a function that pools all flags and executes each accordignly

		// mini functions for each flag

		// error handling function

	}
}
