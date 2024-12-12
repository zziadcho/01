package functions

import (
	"fmt"
	"io/fs"
	"strings"
	"time"
	"unicode"
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
			if slice[j].Time.After(slice[j+1].Time) {
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
		if !unicode.IsDigit(rune(filename[i])) && !unicode.IsLetter(rune(filename[i])) {
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

func FormatTime(t time.Time) string {
	// Format the time normally
	formatted := t.Format("Jan 02 15:04")

	// Extract the day part, ensuring only it is adjusted
	day := t.Day()
	if day < 10 {
		formatted = fmt.Sprintf("%s  %d %s", t.Format("Jan"), day, t.Format("15:04"))
	}

	return formatted
}
