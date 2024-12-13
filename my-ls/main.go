package main

import (
	"fmt"
	"my-ls/functions"
	"os"
	"strings"
)

func main() {
	// Parse flags from command-line arguments
	flags, err := functions.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Handle "Help" flag
	if flags["Help"] {
		fmt.Printf("Usage: myls [OPTION]... [FILE]...\nOptions:\n  -R, --recursive    list directories recursively\n  -r, --reverse      reverse order\n  -a, --all          include hidden files\n  -l                 long listing format\n  -t                 sort by modification time\n")
		return
	}

	// If no paths are provided, default to current directory
	paths := []string{}
	if len(os.Args) == 1 || len(flags) == 0 {
		paths = append(paths, ".")
	} else {
		// Separate out the paths from flags
		for _, arg := range os.Args[1:] {
			if !strings.HasPrefix(arg, "-") {
				paths = append(paths, arg)
			}
		}
	}

	// Sort the paths if necessary (this is already handled by the function)
	functions.SortStringByNam(paths)

	// Handle the "Recursive" flag by printing appropriate headers and recursive calls
	// For each path, call the MyLS function to list directory contents
	for i, path := range paths {
		pathInfo, err := os.Stat(path)
		if err != nil {
			fmt.Printf("myls: cannot access '%v': %v\n", path, err)
			continue
		}

		// If path is a directory, print its contents, otherwise, just print the file
		if pathInfo.IsDir() {
			// If it's a directory and we have multiple paths, print the path name before listing
			if len(paths) > 1 && i > 0 {
				fmt.Printf("\n")
			}

			// Print the directory header
			if len(paths) > 1 {
				fmt.Printf("%v:\n", path)
			}

			// Call MyLS to list the directory contents
			err = functions.MyLS(path, flags, false)
			if err != nil {
				fmt.Printf("myls: cannot access '%v': %v\n", path, err)
			}

			// If recursive flag is set, list subdirectories recursively
			if flags["Recursive"] {
				err = functions.MyLS(path, flags, true)
				if err != nil {
					fmt.Printf("myls: cannot access '%v': %v\n", path, err)
				}
			}

		} else {
			// If it's a regular file, print its details (do not readdir)
			err = functions.MyLS(path, flags, false)
			if err != nil {
				fmt.Printf("myls: cannot access '%v': %v\n", path, err)
			}
		}
	}

	// If there are any regular files or directories, print a newline at the end of their listing
	if len(paths) > 0 && !flags["Recursive"] {
		fmt.Printf("\n")
	}
}
