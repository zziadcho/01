package source

import (
	"fmt"
	"runtime"
	"log"
	"os/exec"
)

func AutoLaunchBrowser(url string) {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	default:
		fmt.Println("Unsupported platform")
		return
	}
	if err := exec.Command(cmd, args...).Start(); err != nil {
		log.Fatal("Failed to open browser:", err)
	}
}