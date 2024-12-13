package functions

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// Helper function to handle symlink-only display
func showSymlink(path string) error {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("error reading symlink: %v", err)
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		linkTarget, err := os.Readlink(path)
		if err != nil {
			return fmt.Errorf("error reading symlink target for %v: %v", path, err)
		}

		// Get user and group info
		uId := strconv.Itoa(int(fileInfo.Sys().(*syscall.Stat_t).Uid))
		gId := strconv.Itoa(int(fileInfo.Sys().(*syscall.Stat_t).Gid))
		if user, err := user.LookupId(uId); err == nil {
			uId = user.Username
		}
		if group, err := user.LookupGroupId(gId); err == nil {
			gId = group.Name
		}
		// Format the output like `ls -l`
		fmt.Printf("lrwxrwxrwx 1 %s %s %d %s %s -> %s",
			uId, gId,
			fileInfo.Size(),
			FormatTime(fileInfo.ModTime()),
			fileInfo.Name(),
			linkTarget)
	}
	return nil
}
