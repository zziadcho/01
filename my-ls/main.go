package main

import (
	"fmt"
	"my-ls/functions"
	"os"
	"strings"
)

func main() {
	flags, err := functions.ParseArgs(os.Args[1:])

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	var paths []string
	var args []string
	for _, arg := range os.Args[1:] {
        // Modified condition to include all non-flag arguments
        if !strings.HasPrefix(arg, "-") || arg == "-" || strings.HasPrefix(arg, "./") {
            paths = append(paths, arg)
            args = append(args, arg)
        }
    }


	functions.SortStringByNam(args)

	// Only print newline after file listing if there are files to list
	if len(args) >= 1 {
		err = functions.MyLS("", flags, false, args)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		// Only print newline if not recursively listing directories afterwards
		if len(paths) == 0 || (len(paths) == 1 && !flags["Recursive"]) {
			fmt.Printf("\n")
		}
	}

	if len(paths) == 0 {
		paths = append(paths, ".")
	}
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if len(paths) == 1 && paths[0] == currentDir && flags["Recursive"] {
		fmt.Printf("%v:\n", currentDir)
	}
	asa7bi := functions.RemoveDuplicates(&paths, args)
	if flags["Help"] {
		fmt.Printf("Usage: myls [OPTION]... [FILE]...\nOptions:\n  -R, --recursive    list directories recursively\n  -r, --reverse      reverse order\n  -a, --all          include hidden files\n  -l                 long listing format\n  -t                 sort by modification time\n")
		return
	}

	for i, path := range paths {
		pathInfos, err := os.Stat(path)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		// Add newline between multiple directory listings
		if pathInfos.IsDir() && len(paths) > 1 && i > 0 {
			fmt.Printf("\n")
		}
		if pathInfos.IsDir() && len(paths)+asa7bi == 1 {
			if flags["Recursive"] {
				err2 := functions.MyLS(path, flags, true, nil)
				if err2 != nil {
					fmt.Printf("myls: cannot access '%v': %v\n", path, err2)
					return
				}
			} else {
				err2 := functions.MyLS(path, flags, false, nil)
				if err2 != nil {
					fmt.Printf("myls: cannot access '%v': %v\n", path, err2)
					return
				}
			}
		} else if pathInfos.IsDir() && len(paths)+asa7bi > 1 {
			err2 := functions.MyLS(path, flags, true, nil)
			if err2 != nil {
				fmt.Printf("myls: cannot access '%v': %v\n", path, err2)
				return
			}
		}
	}
}
