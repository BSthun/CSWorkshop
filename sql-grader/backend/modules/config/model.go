package iconfig

type Config struct {
	Environment uint8  `yaml:"environment" validate:"gte=1,lte=2"`
	LogLevel    uint32 `yaml:"log_level" validate:"required"`

	Address string   `yaml:"address" validate:"required"`
	BaseUrl string   `yaml:"base_url" validate:"url"`
	Cors    []string `yaml:"cors" validate:"required"`

	CoreMySqlDsn   string `yaml:"core_mysql_dsn" validate:"required"`
	BranchMySqlDsn string `yaml:"branch_mysql_dsn" validate:"required"`

	SentryDsn              string  `yaml:"sentry_dsn" validate:"url"`
	SentryTracesSampleRate float64 `yaml:"sentry_traces_sample_rate" validate:"required"`

	JwtSecret string `yaml:"jwt_secret" validate:"required"`
}
