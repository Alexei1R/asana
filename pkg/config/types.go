package config

import "time"

type Config struct {
	Asana AsanaConfig `mapstructure:"asana"`

	Fetch FetchConfig  `mapstructure:"fetch"`
	Cache ChacheConfig `mapstructure:"cache"`

	Refresh RefreshConfig `mapstructure:"refresh"`
}

type AsanaConfig struct {
	BaseURL     string `mapstructure:"base_url"`
	AccessToken string `mapstructure:"token"`
}

type FetchConfig struct {
	ShortInterval      time.Duration `mapstructure:"short_interval"`
	LongInterval       time.Duration `mapstructure:"long_interval"`
	PollInterval       time.Duration `mapstructure:"polling_interval"`
	SecondPollInterval time.Duration `mapstructure:"second_polling_interval"`
}
type ChacheConfig struct {
	Path string `mapstructure:"path"`
}

type RefreshConfig struct {
	Interval time.Duration `mapstructure:"interval"`
	Retry    int           `mapstructure:"retry"`
}
