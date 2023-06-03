package backdrop

import "time"

func GetDateTime() (*string, *string) {
	d := time.Now().Local().Format("January 2, 2006")
	//d := "May 26, 2023"
	t := time.Now().Local().Format("15:04")
	//t := "11:45"
	return &d, &t
}
