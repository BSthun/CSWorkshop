package data

type CityWrapper struct {
	Cities []*City `yaml:"city"`
}

type City struct {
	ID          int    `yaml:"id"`
	Name        string `yaml:"name"`
	StateID     int    `yaml:"state_id"`
	StateCode   string `yaml:"state_code"`
	StateName   string `yaml:"state_name"`
	CountryID   int    `yaml:"country_id"`
	CountryCode string `yaml:"country_code"`
	CountryName string `yaml:"country_name"`
	Latitude    string `yaml:"latitude"`
	Longitude   string `yaml:"longitude"`
	WikiDataID  string `yaml:"wikiDataId"`
}
