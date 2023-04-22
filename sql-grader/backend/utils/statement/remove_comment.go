package statement

import (
	"regexp"
	"strings"
)

func RemoveComments(s string) *string {
	// Remove comments
	var output string
	for len(s) > 0 {
		i := strings.Index(s, "/*")
		if i == -1 {
			output += s
			break
		}
		output += s[:i]
		j := strings.Index(s, "*/")
		if j == -1 {
			break
		}
		s = s[j+2:]
	}

	// Remove extra spaces
	space := regexp.MustCompile(`\s+`)
	output = space.ReplaceAllString(output, " ")

	// Trim spaces
	output = strings.Trim(output, " ")

	// Return output
	return &output
}
