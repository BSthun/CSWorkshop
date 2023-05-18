package track

import (
	"fmt"
	"mock/modules"
	"strings"
)

var trackRef map[string]bool
var albumRef map[string]bool
var artistRef map[string]*uint64

func Run(clean bool) {
	trackRef = make(map[string]bool)
	albumRef = make(map[string]bool)
	artistRef = make(map[string]*uint64)
	for i := 1; i <= 33; i++ {
		dsn := strings.Replace(modules.Conf.SpotifyScraperDsn, "{{DB_NAME}}", fmt.Sprintf("spotify_prod_u%03d", i), 1)
		DbConn(dsn)
	}
}
