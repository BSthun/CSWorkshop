package data

type PodcastData struct {
	RSS RSS `json:"rss"`
}

type RSS struct {
	Channel Channel `json:"channel"`
}

type Channel struct {
	Link          []string `json:"link"`
	LastBuildDate string   `json:"lastBuildDate"`
	Title         string   `json:"title"`
	Generator     string   `json:"generator"`
	Description   string   `json:"description"`
	Subtitle      string   `json:"subtitle"`
	Author        string   `json:"author"`
	Summary       string   `json:"summary"`
	Language      string   `json:"language"`
	Copyright     string   `json:"copyright"`
	Owner         Owner    `json:"owner"`
	Image         []Image  `json:"image"`
	Category      []string `json:"category"`
	Keywords      string   `json:"keywords"`
	Explicit      string   `json:"explicit"`
	Rating        string   `json:"rating"`
	Item          []Item   `json:"item"`
}

type Image struct {
	URL    string `json:"url"`
	Title  string `json:"title"`
	Link   string `json:"link"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Item struct {
	Title       string  `json:"title"`
	Link        string  `json:"link"`
	Comments    string  `json:"comments"`
	PubDate     string  `json:"pubDate"`
	GUID        string  `json:"guid"`
	Description string  `json:"description"`
	Encoded     string  `json:"encoded"`
	Enclosure   string  `json:"enclosure"`
	Subtitle    string  `json:"subtitle"`
	Summary     string  `json:"summary"`
	Author      string  `json:"author"`
	Explicit    string  `json:"explicit"`
	Block       string  `json:"block"`
	Duration    string  `json:"duration"`
	Episode     int64   `json:"episode"`
	EpisodeType string  `json:"episodeType"`
	Content     Content `json:"content"`
	Image       string  `json:"image"`
}

type Owner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Content struct {
	Title string `json:"title"`
}
