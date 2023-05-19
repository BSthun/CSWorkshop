package country

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"math/rand"
	"mock/modules"
	"mock/types/data"
	"mock/types/model"
	"os"
)

func Run(clean bool) {
	cityWrapper := new(data.CityWrapper)

	cityBytes, err := os.ReadFile("./resources/cities.yml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(cityBytes, cityWrapper); err != nil {
		panic(err)
	}

	addedState := make(map[string]bool)
	for _, city := range cityWrapper.Cities {
		name := fmt.Sprintf("%s, %s", city.StateName, city.CountryName)
		if _, ok := addedState[name]; ok {
			continue
		}
		if city.CountryCode != "TH" {
			if rand.Intn(2) > 0 {
				continue
			}
		}
		addedState[name] = true
		country := &model.Country{
			Id:       nil,
			Name:     &name,
			Users:    nil,
			Concerts: nil,
		}
		if result := modules.DB.Create(country); result.Error != nil {
			panic(result.Error)
		}
	}
}
