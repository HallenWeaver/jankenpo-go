package config

import "github.com/spf13/viper"

// Config holds application-wide configurations, except for database connections, which are in `connections`.
// We are trying to be agnostic here, so we don't want to mix concerns.
type Config struct {
	AppPort string
	GoEnv   string
}

// LoadConfig initializes environment variables and checks critical configurations, just the important ones.
func LoadConfig() *Config {
	return &Config{
		AppPort: viper.GetString("app.port"),
		GoEnv:   viper.GetString("app.environment"),
	}
}
