package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type LongFormatInfo struct {
	Permissions fs.FileMode
	Nlink       string
	User        string
	Group       string
	Size        int
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

func MyLS(path string, flags map[string]bool, showPath bool) error {
	if showPath {
		fmt.Printf("%v:\n", path)
	}
	masterSlice := []LongFormatInfo{}
	var totalBlocks int64
	list, err := ReadAll(path)
	if err != nil {
		return err
	}
	var uId, gId, nLink string
	for _, item := range list {
		if stat, ok := item.Sys().(*syscall.Stat_t); ok {
			if flags["All"] {
				totalBlocks += stat.Blocks
			} else if strings.HasPrefix(item.Name(), ".") {

			} else {
				totalBlocks += stat.Blocks
			}
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
		element := LongFormatInfo{item.Mode(), nLink, uId, gId, int(item.Size()), item.ModTime(), item.Name()}
		masterSlice = append(masterSlice, element)

	}
	if flags["Time"] {
		SortMasterSlice(masterSlice)
	}

	if flags["Reverse"] {
		ReverseMasterSlice(masterSlice)
	}
	var maxNlinkLen, maxUserLen, maxGroupLen, maxLenSize, maxLenTime, maxFileNameLen int

	if flags["LongFormat"] && flags["All"] {
		// First, calculate the maximum lengths for each field
		for _, item := range masterSlice {

			if len(item.Nlink) > maxNlinkLen {
				maxNlinkLen = len(item.Nlink)
			}
			if len(item.User) > maxUserLen {
				maxUserLen = len(item.User)
			}
			if len(item.Group) > maxGroupLen {
				maxGroupLen = len(item.Group)
			}
			if len(strconv.Itoa(item.Size)) > maxLenSize {
				maxLenSize = len(strconv.Itoa(item.Size))
			}
			if len(item.Time.Format("Jan 02 15:04")) > maxLenTime {
				maxLenTime = len(item.Time.Format("Jan 02 15:04"))
			}
			if len(item.FileName) > maxFileNameLen {
				maxFileNameLen = len(item.FileName)
			}
		}
	/***************** ERROR NEED WORK ******************/
		// Print total blocks
		fmt.Printf("total %v\n", totalBlocks/2)
	
		// Now, print the formatted output with dynamic widths
		for _, item := range masterSlice {
			// Correct the formatting string to match the dynamic lengths
			fmt.Printf("%-*s %-*s %-*s %-*s %*d %-*s %-*s\n",
				item.Permissions,   // Dynamic width for Permissions
				item.Nlink, maxNlinkLen,         // Dynamic width for Nlink
				item.User, maxUserLen,           // Dynamic width for User
				item.Group, maxGroupLen,         // Dynamic width for Group
				item.Size, maxLenSize,           // Dynamic width for Size
				item.Time.Format("Jan 02 15:04"), maxLenTime, // Dynamic width for Time
				item.FileName, maxFileNameLen,   // Dynamic width for FileName
			)
		}
	} else if !flags["LongFormat"] && flags["All"] {
		for _, item := range masterSlice {
			fmt.Printf("%v  ", item.FileName)
		}
		println()
	} else if flags["LongFormat"] && !flags["All"] {
		fmt.Printf("total %v\n", totalBlocks/2)
		for _, item := range masterSlice {
			if !strings.HasPrefix(item.FileName, ".") {
				fmt.Printf("%v %"+strconv.Itoa(len(item.Nlink))+"s %-5s %-5s %7d %-10s %s\n",
					item.Permissions,
					item.Nlink,
					item.User,
					item.Group,
					item.Size,
					item.Time.Format("Jan 02 15:04"),
					item.FileName,
				)
			}
		}
	} else {
		for _, item := range masterSlice {
			if !strings.HasPrefix(item.FileName, ".") {
				fmt.Printf("%v  ", item.FileName)
			}
		}
		println()
	}
	for _, item := range list {
		if flags["Recursive"] && item.IsDir() && !strings.HasPrefix(item.Name(), ".") {
			fmt.Printf("\n")
			newPath := path + "/" + item.Name()
			err := MyLS(newPath, flags, true)
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
		fmt.Printf("%v\n", err)
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
		fmt.Printf("Usage: myls [OPTION]... [FILE]...\nOptions:\n  -R, --recursive    list directories recursively\n  -r, --reverse      reverse order\n  -a, --all          include hidden files\n  -l                 long listing format\n  -t                 sort by modification time\n")
		return
	}

	for i, path := range paths {
		pathInfos, err := os.Stat(path)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		if pathInfos.IsDir() && len(paths) == 1 {
			err2 := MyLS(path, flags, false)
			if err2 != nil {
				fmt.Printf("myls: cannot access '%v': %v\n", path, err2)
			}
		} else if pathInfos.IsDir() && len(paths) > 1 {
			err2 := MyLS(path, flags, true)
			if err2 != nil {
				fmt.Printf("myls: cannot access '%v': %v\n", path, err2)
			}
		} else {
			fmt.Printf("%v  \n", path)
		}
		if i != len(paths)-1 {
			fmt.Printf("\n")
		}
	}
}
