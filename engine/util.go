package engine

import (
	"fmt"
	"strings"
)

func ifPXexists(commandParts []string) (string, int, bool) {
	for i := 0; i < len(commandParts)-1; i++ {
		if strings.ToUpper(commandParts[i]) == "PX" && commandParts[i+1] != "" {
			return commandParts[i+1], i, true
		}
	}
	return "", 0, false
}

func PrintHelp() {
	fmt.Println("Own Redis")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  own-redis [--port <N>]")
	fmt.Println("  own-redis --help")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help       Show this screen.")
	fmt.Println("  --port N     Port number.")
}
