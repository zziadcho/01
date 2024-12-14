package functions

import (
	"fmt"
	"io/fs"
	"strings"
	"time"
)

type LongFormatInfo struct {
	Permissions fs.FileMode
	Nlink       string
	User        string
	Group       string
	Major       string
	Minor       string
	Size        int
	Time        time.Time
	FileName    string
}

func SortByTime(slice []LongFormatInfo) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j].Time.Before(slice[j+1].Time) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func SortLs(slice []LongFormatInfo) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if strings.ToLower(getKey(slice[i].FileName)) > strings.ToLower(getKey(slice[j].FileName)) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
func getKey(filename string) string {
	for i := 0; i < len(filename); i++ {
		if !IsDigit(rune(filename[i])) && !IsLetter(rune(filename[i])) {
			filename = filename[:i] + filename[i+1:]
			i--
		}
	}
	return filename
}

func SortByName(slice []LongFormatInfo) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if strings.HasPrefix(slice[j].FileName, ".") {
				slice[j].FileName = strings.Trim(slice[j].FileName, ".")
			}
			if strings.HasPrefix(slice[j+1].FileName, ".") {
				slice[j+1].FileName = strings.Trim(slice[j+1].FileName, ".")
			}
			if slice[j].FileName > slice[j+1].FileName {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func SortStringByNam(slice []string) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func ReverseOrder(slice []LongFormatInfo) []LongFormatInfo {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func FormatTime(z time.Time) string {
	a, b, c, d, res := z.Month(), z.Day(), z.Year(), fmt.Sprintf("%02d:%02d", z.Hour(), z.Minute()), ""
	ok := time.Now().Sub(z)
	ko := ok.Hours()
	if ko < 4380 {
		res = fmt.Sprintf("%s %2d %5s", fmt.Sprintf("%v", a)[:3], b, d)
	} else {
		res = fmt.Sprintf("%s %2d %5d", fmt.Sprintf("%v", a)[:3], b, c)
	}
	return res
}

func IsLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func IsDigit(r rune) bool {
	return (r >= '0' && r <= '9')
}

func RemoveDuplicates(slice1 *[]string, slice2 []string) int {
	duplicates := make(map[string]bool)
	for _, elem := range slice2 {
		duplicates[elem] = true
	}

	original := *slice1
	filtered := original[:0]
	for _, elem := range original {
		if !duplicates[elem] {
			filtered = append(filtered, elem)
		}
	}
	asa7bi := len(*slice1)
	*slice1 = filtered

	return asa7bi - len(filtered)
}
