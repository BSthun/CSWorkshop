package statement

import (
	"strings"
)

func ParseDatabaseName(stmt string) string {
	// Split the statement by whitespace characters
	parts := strings.Fields(stmt)
	for i := range parts {
		part := strings.ToLower(parts[i])
		if part == "from" && i+1 < len(parts) {
			// The next part should contain the table name or alias,
			// which can include the database name as a prefix.
			table := parts[i+1]
			if idx := strings.IndexByte(table, '.'); idx >= 0 {
				return strings.Trim(table[:idx], "`")
			}
			// If the table name does not include a prefix with a dot,
			// we assume that it belongs to the default database.
			return "default"
		} else if part == "into" && i+1 < len(parts) {
			// The next part should contain the table name.
			table := parts[i+1]
			if idx := strings.IndexByte(table, '.'); idx >= 0 {
				return strings.Trim(table[:idx], "`")
			}
			// If the table name does not include a prefix with a dot,
			// we assume that it belongs to the default database.
			return "default"
		} else if part == "update" && i+1 < len(parts) {
			// The next part should contain the table name.
			table := parts[i+1]
			if idx := strings.IndexByte(table, '.'); idx >= 0 {
				return strings.Trim(table[:idx], "`")
			}
			// If the table name does not include a prefix with a dot,
			// we assume that it belongs to the default database.
			return "default"
		} else if part == "delete" && i+1 < len(parts) {
			// The next part should be "from", followed by the table name.
			if i+2 < len(parts) && strings.ToLower(parts[i+1]) == "from" {
				table := parts[i+2]
				if idx := strings.IndexByte(table, '.'); idx >= 0 {
					return strings.Trim(table[:idx], "`")
				}
				// If the table name does not include a prefix with a dot,
				// we assume that it belongs to the default database.
				return "default"
			}
			// If the statement is not in the expected format, return an empty string.
			return ""
		}
	}
	return ""
}
