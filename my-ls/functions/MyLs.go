package functions

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

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
	var accumulatedLength int
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
		accumulatedLength += len(item.Name())
		masterSlice = append(masterSlice, element)
	}
	if flags["Time"] {
		SortByTime(masterSlice)
	} else {
		SortByName(masterSlice)
	}

	if flags["Reverse"] {
		ReverseOrder(masterSlice)
	}
	var maxPermLen, maxNlinkLen, maxUserLen, maxGroupLen, maxLenSize, maxLenTime, maxFileNameLen int

	if flags["LongFormat"] && flags["All"] {
		for _, item := range masterSlice {
			permStr := strconv.Itoa(int(item.Permissions))
			if len(permStr) > maxPermLen {
				maxPermLen = len(permStr)
			}

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

		fmt.Printf("total %v\n", totalBlocks/2)

		for _, item := range masterSlice {
			symLinkArrow := ""
			if item.Permissions&os.ModeSymlink != 0 {
				linkTarget, err := os.Readlink(path + "/" + item.FileName)
				if err == nil {
					symLinkArrow = fmt.Sprintf(" -> %s", linkTarget)
				}
			}

			fmt.Printf("%-*s %*s %-*s %-*s %*d %-*s %s\n",
				maxPermLen, strings.ToLower(fmt.Sprintf("%v ", item.Permissions)),
				maxNlinkLen, item.Nlink,
				maxUserLen, item.User,
				maxGroupLen, item.Group,
				maxLenSize, item.Size,
				maxLenTime, FormatTime(item.Time),
				item.FileName+symLinkArrow,
			)
		}

	} else if !flags["LongFormat"] && flags["All"] {
		for _, item := range masterSlice {
			fmt.Printf("%v  ", item.FileName)
		}
		fmt.Printf("\n")
	} else if flags["LongFormat"] && !flags["All"] {
		for _, item := range masterSlice {
			permStr := strconv.Itoa(int(item.Permissions))
			if len(permStr) > maxPermLen {
				maxPermLen = len(permStr)
			}

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

		fmt.Printf("total %v\n", totalBlocks/2)
		for _, item := range masterSlice {
			if !strings.HasPrefix(item.FileName, ".") {
				symLinkArrow := ""
				if item.Permissions&os.ModeSymlink != 0 {
					linkTarget, err := os.Readlink(path + "/" + item.FileName)
					if err == nil {
						symLinkArrow = fmt.Sprintf(" -> %s", linkTarget)
					}
				}

				fmt.Printf("%-*s %*s %-*s %-*s %*d %-*s %s\n",
					maxPermLen, strings.ToLower(fmt.Sprintf("%v ", item.Permissions)),
					maxNlinkLen, item.Nlink,
					maxUserLen, item.User,
					maxGroupLen, item.Group,
					maxLenSize, item.Size,
					maxLenTime, FormatTime(item.Time),
					item.FileName+symLinkArrow,
				)
			}

		}
	} else if 2*(len(masterSlice)-1)+accumulatedLength < 132 {
		for _, item := range masterSlice {
			if !strings.HasPrefix(item.FileName, ".") {
				fmt.Printf("%v  ", item.FileName)
			}
		}
		fmt.Printf("\n")
	} else {
		for _, item := range masterSlice {
			if !strings.HasPrefix(item.FileName, ".") {
				fmt.Printf("%v\n", item.FileName)
			}
		}
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
