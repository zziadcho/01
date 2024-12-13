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
	masterSlice := []LongFormatInfo{}
	var totalBlocks int64
	var uId, gId, nLink, major, minor string
	var accumulatedLength int
	list, err := ReadAll(path)
	if err != nil {
		return err
	}
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
			if stat.Mode&syscall.S_IFBLK != 0 || stat.Mode&syscall.S_IFCHR != 0 {
				major, minor = fmt.Sprintf("%v,", Major(stat.Rdev)), fmt.Sprintf("%v", Minor(stat.Rdev))
			} else {
				major = "0"
				minor = "0"
			}
		}
		if user, err := user.LookupId(uId); err == nil {
			uId = user.Username
		}
		if group, err := user.LookupGroupId(gId); err == nil {
			gId = group.Name
		}
		element := LongFormatInfo{item.Mode(), nLink, uId, gId, major, minor, int(item.Size()), item.ModTime(), item.Name()}
		accumulatedLength += len(item.Name())
		masterSlice = append(masterSlice, element)
		//fmt.Println(major)
	}
	if flags["Time"] {
		SortByTime(masterSlice)
	} else {
		SortLs(masterSlice)
	}

	if flags["Reverse"] {
		ReverseOrder(masterSlice)
	}
	var maxPermLen, maxNlinkLen, maxUserLen, maxGroupLen, maxLenSize, maxLenTime, maxFileNameLen, maxMinorLen, maxMajorLen int

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
			if len(item.Minor) > maxMinorLen && item.Major != "0," {
				maxMinorLen = len(item.Minor)
			}

			if len(item.Major) > maxMajorLen && item.Minor != "0" {
				maxMajorLen = len(item.Major)
			}
		}

		fmt.Printf("total %v\n", totalBlocks/2)

		for _, item := range masterSlice {
			symLinkArrow := ""
			if item.Permissions&os.ModeSymlink != 0 {
				linkTarget, err := os.Readlink(path + "/" + item.FileName)
				if err == nil {
					symLinkArrow = fmt.Sprintf("-> %s", linkTarget)
				}
			}
			formattedPerms := formatPermissions(item.Permissions)
			placeHolder := ""
			if item.Permissions&os.ModeDevice != 0 || item.Permissions&os.ModeCharDevice != 0 {
				fmt.Printf("%*s %*s %-*s %-*s %*s %*s %-*s %s %s\n",
					maxPermLen, formattedPerms,
					maxNlinkLen, item.Nlink,
					maxUserLen, item.User,
					maxGroupLen, item.Group,
					maxMajorLen, item.Major,
					maxLenSize, item.Minor,
					maxLenTime, FormatTime(item.Time),
					item.FileName,
					symLinkArrow,
				)
			} else {
				fmt.Printf("%*s %*s %-*s %-*s %*s %*d %-*s %s %s\n",
					maxPermLen, formattedPerms,
					maxNlinkLen, item.Nlink,
					maxUserLen, item.User,
					maxGroupLen, item.Group,
					maxMajorLen, placeHolder,
					maxLenSize, item.Size,
					maxLenTime, FormatTime(item.Time),
					item.FileName,
					symLinkArrow,
				)
			}
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

			if len(item.Minor) > maxMinorLen && item.Major != "0," {
				maxMinorLen = len(item.Minor)
			}

			if len(item.Major) > maxMajorLen && item.Minor != "0" {
				maxMajorLen = len(item.Major)
			}
		}

		fmt.Printf("total %v\n", totalBlocks/2)

		for _, item := range masterSlice {
			if strings.HasPrefix(item.FileName, ".") {
				continue
			}
			symLinkArrow := ""
			if item.Permissions&os.ModeSymlink != 0 {
				linkTarget, err := os.Readlink(path + "/" + item.FileName)
				if err == nil {
					symLinkArrow = fmt.Sprintf("-> %s", linkTarget)
				}
			}
			formattedPerms := formatPermissions(item.Permissions)
			if item.Permissions&os.ModeDevice != 0 || item.Permissions&os.ModeCharDevice != 0 {
				fmt.Printf("%*s %*s %-*s %-*s %*s %*s %-*s %s %s\n",
					maxPermLen, formattedPerms,
					maxNlinkLen, item.Nlink,
					maxUserLen, item.User,
					maxGroupLen, item.Group,
					maxMajorLen, item.Major,
					maxLenSize, item.Minor,
					maxLenTime, FormatTime(item.Time),
					item.FileName,
					symLinkArrow,
				)
			} else {
				fmt.Printf("%*s %*s %*s %*s %*d %-*s %s %s\n",
					maxPermLen, formattedPerms,
					maxNlinkLen, item.Nlink,
					maxUserLen, item.User,
					maxGroupLen, item.Group,
					maxLenSize, item.Size,
					maxLenTime, FormatTime(item.Time),
					item.FileName,
					symLinkArrow,
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

	for _, item := range masterSlice {
		if flags["Recursive"] && item.Permissions.IsDir() && item.FileName != "." && item.FileName != ".." {
			if !flags["All"] && strings.HasPrefix(item.FileName, ".") {
				continue
			}
			fmt.Printf("\n")
			newPath := path + "/" + item.FileName
			err := MyLS(newPath, flags, true)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
