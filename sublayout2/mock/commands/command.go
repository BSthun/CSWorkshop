package commands

import (
	"flag"
	"fmt"
	"mock/commands/podcast"
	"os"
)

func Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Mock manager requires subcommand\n")
		return
	}

	clean := flag.Bool("clean", false, "a bool")

	switch args[1] {
	case "podcast":
		podcast.Run(*clean)
	default:
		fmt.Printf("Unknown subcommand\n")
	}
}
