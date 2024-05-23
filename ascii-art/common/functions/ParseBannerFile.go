package functions

func ParseBannerFile(arg string) bool {
	if arg == "standard" || arg == "shadow" || arg == "thinkertoy" {
		return true
	}
	return false
}
