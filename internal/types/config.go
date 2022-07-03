package types

type Dashboard struct {
	Title      string
	Categories []Category
}

type Category struct {
	Name, Icon string
	Entries    []Entry
}

type Entry struct {
	Name            string
	Url             string
	Icon            string
	IconURL         string
	HealthCheckPath string
	UseNewTab       bool
}

type EnvConfig struct {
	DashBoarConfigPath       string `mapstructure:"DASHBOAR_CONFIG_FILE"`
	Port                     int    `mapstructure:"DASHBOAR_PORT"`
	LimiterMaxRequests       int    `mapstructure:"LIMITER_MAX_REQUESTS"`
	LimiterExpireTimeSeconds int    `mapstructure:"LIMITER_EXPIRE_TIME_SECONDS"`
	CacheExpireTimeSeconds   int    `mapstructure:"CACHE_EXPIRE_TIME_SECONDS"`
}
