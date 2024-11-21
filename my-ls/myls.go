package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strings"
	"syscall"
	"time"
)

var currentDir = "."

type LongFormatInfo struct {
	Permissions fs.FileMode
	Nlink       string
	User        string
	Group       string
	Size        string
	Time        time.Time
	FileName    string
}

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
	}
	return Flags, nil
}

func humanReadableSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "KMGTPE"[exp])
}

func SortMasterSlice(slice []LongFormatInfo) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			// Compare ModTime (earlier to later)
			if slice[j].Time.After(slice[j+1].Time) {
				// Swap if needed
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func ReverseMasterSlice(slice []LongFormatInfo) []LongFormatInfo {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func MyLS(path string, flags map[string]bool) error {
	masterSlice := []LongFormatInfo{}
	list, err := ReadAll(path)
	if err != nil {
		return err
	}
	var uId, gId, nLink string
	for _, item := range list {
		if !flags["All"] && !strings.HasPrefix(item.Name(), ".") {
			if stat, ok := item.Sys().(*syscall.Stat_t); ok {
				uId = fmt.Sprintf("%d", stat.Uid)
				gId = fmt.Sprintf("%d", stat.Gid)
				nLink = fmt.Sprintf("%d", stat.Nlink)
			}
			if user, err := user.LookupId(uId); err == nil {
				uId = user.Username
			}
			if group, err := user.LookupGroupId(gId); err == nil {
				gId = group.Name
			}
			element := LongFormatInfo{item.Mode(), nLink, uId, gId, humanReadableSize(item.Size()), item.ModTime(), item.Name()}
			masterSlice = append(masterSlice, element)
		}
	}
	if flags["Time"] {
		SortMasterSlice(masterSlice)
	}

	if flags["Reverse"] {
		ReverseMasterSlice(masterSlice)
	}
	if flags["LongFormat"] {
		for _, item := range masterSlice {
			fmt.Printf("%v %2s %-5s %-5s %5s %-10s %-10s\n",
				item.Permissions,
				item.Nlink,
				item.User,
				item.Group,
				item.Size,
				item.Time.Format("Jan 02 15:04"),
				item.FileName,
			)
		}
	} else {
		for _, item := range masterSlice {
			fmt.Printf("%v  ", item.FileName) // complete behavior for --all flag activated
		}
		println()
	}
	for _, item := range list {
		if flags["Recursive"] && item.IsDir() && !strings.HasPrefix(item.Name(), ".") {
			println()
			newPath := path + "/" + item.Name()
			fmt.Printf("%v:\n", newPath)
			err := MyLS(newPath, flags)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
