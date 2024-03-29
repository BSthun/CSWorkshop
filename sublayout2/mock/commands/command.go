package commands

import (
	"flag"
	"fmt"
	"mock/commands/country"
	"mock/commands/podcast"
	"mock/commands/track"
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
	case "country":
		country.Run(*clean)
	case "track":
		track.Run(*clean)
	default:
		fmt.Printf("Unknown subcommand\n")
	}
}
