package functions

import (
	"io/fs"
	"os"
)

func ReadAll(path string) ([]fs.FileInfo, error) {
	var List []fs.FileInfo

	items, err := os.ReadDir(path)
	if err != nil {
		return List, err
	}

	currentDir, err := os.Stat(path + "/" + ".")
	if err != nil {
		return List, err
	}
	parentDir, err := os.Stat(path + "/" + "..")
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
