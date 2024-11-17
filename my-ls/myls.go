package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ReadAll(path string) ([]fs.FileInfo, error) {
	var List []fs.FileInfo

	items, err := os.ReadDir(path)
	if err != nil {
		return List, errors.New("error reading the directory: " + err.Error())
	}

	currentDir, err := os.Stat(".")
	if err != nil {
		return List, err
	}
	parentDir, err := os.Stat("..")
	if err != nil {
		return List, err
	}

	List = append(List, currentDir, parentDir)

	for _, item := range items {
		itemInfo, err := item.Info()
		if err != nil {
			return List, err
		}
		List = append(List, itemInfo)
	}
	return List, nil
}

func ParseArgs(args []string) (map[string]bool, error) {
	Flags := make(map[string]bool)

	Flags["LongFormat"] = false
	Flags["Recursive"] = false
	Flags["Reverse"] = false
	Flags["Time"] = false
	Flags["Help"] = false
	Flags["All"] = false

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			arg = strings.TrimPrefix(arg, "--")

			for i := 0; i < len(arg); i++ {
				switch arg {
				case "recursive":
					Flags["Recursive"] = true
				case "reverse":
					Flags["Reverse"] = true
				case "all":
					Flags["All"] = true
				case "help":
					Flags["Help"] = true
				default:
					return Flags, fmt.Errorf("myls: unrecognized option -- '%v'\nTry 'myls --help' for more information", string(arg[i]))
				}
			}
		} else if strings.HasPrefix(arg, "-") {
			arg = strings.TrimPrefix(arg, "-")

			for i := 0; i < len(arg); i++ {
				switch arg[i] {
				case 'R':
					Flags["Recursive"] = true
				case 'r':
					Flags["Reverse"] = true
				case 'a':
					Flags["All"] = true
				case 't':
					Flags["Time"] = true
				case 'l':
					Flags["LongFormat"] = true
				default:
					return Flags, fmt.Errorf("myls: invalid option -- '%v'\nTry 'myls --help' for more information", string(arg[i]))
				}
			}
		}
		// if strings.HasPrefix(arg, "--") {
		// 	if arg == "--help" {
		// 		Flags["Help"] = true
		// 	} else {
		// 		return Flags, fmt.Errorf("myls: unknown option '%s'. Try './myls --help' for more information", arg)
		// 	}
		// } else if strings.HasPrefix(arg, "-") {
		// 	for i := 1; i < len(arg); i++ {
		// 		if !strings.ContainsAny(string(arg[i]), validFlags) {
		// 			return Flags, fmt.Errorf("myls: unknown option -%v. Try './myls --help' for more information", string(arg[i]))
		// 		}
		// 		switch arg[i] {
		// 		case 'l':
		// 			Flags["LongFormat"] = true
		// 		case 'R':
		// 			Flags["Recursive"] = true
		// 		case 'r':
		// 			Flags["Reverse"] = true
		// 		case 't':
		// 			Flags["Time"] = true
		// 		case 'a':
		// 			Flags["All"] = true
		// 		case 'h':
		// 			Flags["Help"] = true
		// 		}
		// 	}
		// } else {
		// 	return Flags, fmt.Errorf("myls: cannot access '%v': No such file or directory", arg)
		// }
	}
	return Flags, nil
}

func main() {
	flags, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	if flags["Help"] {
		fmt.Println("Usage: myls [OPTION]... [FILE]...\nOptions:\n  -R, --recursive    list directories recursively\n  -r, --reverse      reverse order\n  -a, --all          include hidden files\n  -l                 long listing format\n  -t                 sort by modification time")
		return
	} 

	for _, path := range paths {
		err := MyLS(path, flags)
		if err != nil {
			fmt.Printf("myls: cannot access '%v': %v\n", path, err)
		}
	}
}


func MyLS(path string, flags map[string]bool) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		items, err := ReadAll(path)
		if err != nil {
			return err
		}

		fmt.Printf("Directory: %s\n", path)
		for _, item := range items {
			if !flags["All"] && strings.HasPrefix(item.Name(), ".") {
				continue // rani 3ad khdam 3la functionality dyal flags
			}
			fmt.Println(item.Name())
		}
	} else {
		fmt.Println(info.Name())
	}

	return nil
}
