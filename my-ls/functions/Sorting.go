package functions

import (
	"fmt"
	"io/fs"
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

func SortByTime(slice []LongFormatInfo) {
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

func SortByName(slice []LongFormatInfo) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
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
