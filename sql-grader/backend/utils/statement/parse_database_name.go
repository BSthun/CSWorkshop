package statement

import (
	"strings"

	"backend/utils/value"
)

func ParseDatabaseName(stmt string) *string {
	// Split the statement by whitespace characters
	parts := strings.Fields(stmt)
	for i, part := range parts {
		part = strings.ToLower(part)
		if (part == "from" || part == "into" || part == "table" || part == "on" || part == "update") && i+1 < len(parts) {
			// The next part should contain the table name or alias,
			// which can include the database name as a prefix.
			table := parts[i+1]
			if idx := strings.IndexByte(table, '.'); idx >= 0 {
				return value.Ptr(strings.Trim(table[:idx], "`"))
			}
			// If the table name does not include a prefix with a dot,
			// we assume that it belongs to the default database.
			return nil
		} else if part == "delete" && i+1 < len(parts) {
			// The next part should be "from", followed by the table name.
			if i+2 < len(parts) && strings.ToLower(parts[i+1]) == "from" {
				table := parts[i+2]
				if idx := strings.IndexByte(table, '.'); idx >= 0 {
					return value.Ptr(strings.Trim(table[:idx], "`"))
				}
				// If the table name does not include a prefix with a dot,
				// we assume that it belongs to the default database.
				return nil
			}
			// If the statement is not in the expected format, return nil.
			return nil
		}
	}
	return nil
}
