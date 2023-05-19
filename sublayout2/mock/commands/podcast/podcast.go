package podcast

import (
	"fmt"
	"math/rand"
	"mock/modules"
	"mock/types/model"
	"mock/utils/data"
	"time"
)

func Run(clean bool) {
	fmt.Printf("podcast command with clean: %v\n", clean)

	data.ReadPodcastData()

	var podcastCategories = [5]string{"Science & Technology", "Business & Management", "Philosophy", "Entertainment", "Design"}

	// Create podcast section first
	for _, category := range podcastCategories {
		var podcastSection *model.PodcastSection

		if result := modules.DB.First(&podcastSection, "name = ?", category); result.RowsAffected == 0 {
			podcastSectionCreate := model.PodcastSection{
				Name: &category,
			}
			if result := modules.DB.Create(&podcastSectionCreate); result.Error != nil {
				panic(result.Error)
			}
		}
	}

	for _, podcast := range data.CurrentData {

		podcastShows := model.PodcastShow{
			Title:       &podcast.RSS.Channel.Title,
			Author:      &podcast.RSS.Channel.Author,
			Description: &podcast.RSS.Channel.Description,
		}

		if result := modules.DB.Create(&podcastShows); result.Error != nil {
			panic(result)
		}

		for _, item := range podcast.RSS.Channel.Item {
			podcastItem := model.PodcastEpisode{
				ShowId:      podcastShows.Id,
				Title:       &item.Title,
				Description: &item.Description,
			}

			if result := modules.DB.Create(&podcastItem); result.Error != nil {
				panic(result.Error)
			}
		}

		for _, category := range podcast.RSS.Channel.Category {
			var podcastCategory *model.PodcastCategory

			if result := modules.DB.First(&podcastCategory, "name = ?", category); result.RowsAffected == 0 {

				var podcastSection *model.PodcastSection

				// Find section
				if result := modules.DB.First(&podcastSection, "name LIKE ?", "%"+category+"%"); result.Error != nil {
					panic(result.Error)
				}

				coverUrl := "https://google.com"

				var featured bool

				rand.Seed(time.Now().UnixNano())

				if rand.Intn(2) == 1 {
					featured = true
				} else {
					featured = false
				}

				podcastCategoryCreate := model.PodcastCategory{
					SectionId: podcastSection.Id,
					Name:      &category,
					CoverUrl:  &coverUrl,
					Featured:  &featured,
				}

				if result := modules.DB.Create(&podcastCategoryCreate); result.Error != nil {
					panic(result.Error)
				}

				podcastShowCategory := model.PodcastShowCategory{
					ShowId:     podcastShows.Id,
					CategoryId: podcastCategoryCreate.Id,
				}

				if result := modules.DB.Create(&podcastShowCategory); result.Error != nil {
					panic(result.Error)
				}
			}
		}
	}

	fmt.Println("Added podcast data successfully")
}
