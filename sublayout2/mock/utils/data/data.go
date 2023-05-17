package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mock/types/data"
	"os"
	"path/filepath"
)

var CurrentData []*data.PodcastData

func ReadPodcastData() {

	// File name
	fileArray := [5]string{"ben_lionel_scott", "nopadol_story", "nuenglc", "ted", "the_standard"}

	for _, fileName := range fileArray {
		// Open our jsonFile
		absPath, _ := filepath.Abs("./utils/data/" + fileName + ".json")

		jsonFile, err := os.Open(absPath)

		// if we os.Open returns an error then handle it
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully Opened " + fileName + ".json")

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result *data.PodcastData
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			log.Fatal(err.Error())
		}

		CurrentData = append(CurrentData, result)
	}

	fmt.Println("Successfully Changed random data")
}
