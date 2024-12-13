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
		if !strings.HasPrefix(arg, "-") || arg == "-" {
			paths = append(paths, arg)
			argInfo, err := os.Stat(arg)
			if err != nil {
				fmt.Println(err)
				return
			}
			if !argInfo.IsDir() || argInfo.Mode()&os.ModeSymlink != 0 {
				args = append(args, arg)
			}
		}
	}
	functions.SortStringByNam(args)

	if len(args) >= 1 {
		err = functions.MyLS("", flags, false, args)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		fmt.Printf("\n")
	}

	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	if flags["Help"] {
		fmt.Printf("Usage: myls [OPTION]... [FILE]...\nOptions:\n  -R, --recursive    list directories recursively\n  -r, --reverse      reverse order\n  -a, --all          include hidden files\n  -l                 long listing format\n  -t                 sort by modification time\n")
		return
	}
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if flags["Recursive"] {
		for _, arg := range args {
			if arg == currentDir {
				fmt.Printf("%v:\n", currentDir)
			} else {
				fmt.Printf(".:\n")
			}
		}
	}
	for i, path := range paths {
		pathInfos, err := os.Stat(path)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		if pathInfos.IsDir() && len(paths) > 1 && i > 0 {
			fmt.Printf("\n")
		}
		if pathInfos.IsDir() && len(paths) == 1 {
			err2 := functions.MyLS(path, flags, false, nil)
			if err2 != nil {
				fmt.Printf("myls: cannot access '%v': %v\n", path, err2)
				return
			}
		} else if pathInfos.IsDir() && len(paths) > 1 {
			err2 := functions.MyLS(path, flags, true, nil)
			if err2 != nil {
				fmt.Printf("myls: cannot access '%v': %v\n", path, err2)
				return
			}
		}
	}
}
