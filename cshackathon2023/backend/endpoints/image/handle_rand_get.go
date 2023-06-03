package image

import (
	"backend/utils/text"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
)

func RandGetHandler(c *fiber.Ctx) error {
	// * Get all files in the directory
	contents, err := ioutil.ReadDir(text.RelativePath("resources/image"))
	if err != nil {
		logrus.Warn(err)
	}

	// * Filter out directories
	var files []fs.FileInfo
	for _, content := range contents {
		if !content.IsDir() {
			files = append(files, content)
		}
	}
	//
	//// * Sort the files by date modified
	//sort.Slice(files, func(i, j int) bool {
	//	return files[i].ModTime().After(files[j].ModTime())
	//})

	// * Calculate skip index
	//last := 300
	//if len(files) < last {
	//	last = len(files)
	//}

	// * Get last 10 files
	//files = files[:last]

	// * Random file in the last 10 files
	file := files[text.Rand.Intn(len(files))]

	// * Redirect to the file
	return c.Redirect("/image/" + file.Name())
}
