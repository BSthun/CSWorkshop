package podcast

import (
	"fmt"
	"mock/utils/data"
)

func Run(clean bool) {
	fmt.Printf("podcast command with clean: %v\n", clean)

	data.ReadPodcastData()
}
