package functions

import (
	"fmt"
	"strings"
)

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
		} else if strings.HasPrefix(arg, "-") && arg != "-" {
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
