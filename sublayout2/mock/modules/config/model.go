package iconfig

type Config struct {
	Environment uint8  `yaml:"environment" validate:"gte=1,lte=2"`
	LogLevel    uint32 `yaml:"log_level" validate:"required"`

	MysqlDsn     string `yaml:"mysql_dsn" validate:"required"`
	MysqlMigrate bool   `yaml:"mysql_migrate" validate:"omitempty"`

	SpotifyScraperDsn string `yaml:"spotify_scraper_dsn" validate:"required"`
}
