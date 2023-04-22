package statement

import "strings"

func ParseUserHostLog(s string) (*string, *string) {
	// Split the input string by '[' and ']' characters
	parts := strings.Split(s, "[")
	if len(parts) != 3 {
		return nil, nil
	}

	// Extract the username from the first part
	username := strings.TrimSpace(parts[0])
	if username == "" {
		return nil, nil
	}

	// Extract the IP address from the third part
	last := strings.Split(parts[2], "]")
	ip := strings.TrimSpace(last[0])
	if ip == "" {
		return nil, nil
	}

	return &username, &ip
}
