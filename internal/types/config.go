package types

type Dashboard struct {
	Title      string     `yaml:"title"`
	Categories []Category `yaml:"categories"`
}

type Category struct {
	Name    string `yaml:"name"`
	Icon    string `yaml:"icon"`
	Entries []Entry
}

type Entry struct {
	Name            string `yaml:"name"`
	Url             string `yaml:"url"`
	Icon            string `yaml:"icon"`
	IconURL         string `yaml:"icon-url"`
	HealthCheckPath string `yaml:"health-check"`
	UseNewTab       bool   `yaml:"use-new-tab"`
}

type EnvConfig struct {
	DashBoarConfigPath       string `mapstructure:"DASHBOAR_CONFIG_FILE"`
	Port                     int    `mapstructure:"DASHBOAR_PORT"`
	LimiterMaxRequests       int    `mapstructure:"LIMITER_MAX_REQUESTS"`
	LimiterExpireTimeSeconds int    `mapstructure:"LIMITER_EXPIRE_TIME_SECONDS"`
	CacheExpireTimeSeconds   int    `mapstructure:"CACHE_EXPIRE_TIME_SECONDS"`
}
