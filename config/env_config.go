package config

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

func SetupEnv() error {
	// Set the file name of the configurations file
	viper.SetConfigName("env")
	// Set the file extension
	viper.SetConfigType("yaml")
	// Set the path to look for the configurations file
	viper.AddConfigPath("environments")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		slog.Error(err.Error())
		slog.Warn("env_{environment}.yaml file not found or failed to load")
		return err
	}

	setViperDefaults()
	slog.Info(fmt.Sprintf("Environment variables loaded successfully, running in %v mode", viper.GetString("app.environment")))

	return nil
}

func setViperDefaults() {
	// Set default value for environment
	viper.SetDefault("app.environment", "development")
}
