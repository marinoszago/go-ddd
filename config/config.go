// Package config contains the necessary app-configuration
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config exported.
type Config struct {
	Server        ServerConfigurations
	Database      DatabaseConfigurations
	ImportService ImportServiceConfigurations
}

// ServerConfigurations exported.
type ServerConfigurations struct {
	Protocol string
	Host     string
	Port     int
}

// DatabaseConfigurations exported.
type DatabaseConfigurations struct {
	Name             string
	User             string
	Password         string
	Host             string
	Port             string
	ConnectionString string
}

// ImportServiceConfigurations is a struct containing the import configurations.
type ImportServiceConfigurations struct {
	Enabled bool
	File    string
}

// CanRun determines if the import will run based on configurations.
func (isc ImportServiceConfigurations) CanRun() bool {
	return isc.Enabled && isc.File != ""
}

// New creates a new config.
func New() (*Config, error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, fmt.Errorf("[err] reading config file, %w", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		return &Config{}, fmt.Errorf("[err] decode to struct, %w", err)
	}

	configuration.Database.ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		configuration.Database.User,
		configuration.Database.Password,
		configuration.Database.Host,
		configuration.Database.Port)

	return &configuration, nil
}
