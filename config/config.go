// Package config provides functionality to manage configuration for the shortofti.me application
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Conf is guideline for allowing alternative configuration implementations
type Conf interface {
	GetListener() string
	GetRootDir() string
	IsDebug() bool
	LoadFromFile() (Conf, error)
}

// Config is the parent structure for all configuration provided data
type Config struct {
	ConfigFile string   `json:"-"`
	Database   Database `json:"database"`
	Debug      bool     `json:"debug"`
	Listener   string   `json:"listener"`
	RootDir    string   `json:"rootDir"`
}

// Database provides a sub-struct just for database connection details
type Database struct {
	Host string `json:"host"`
	Name string `json:"name"`
	Pass string `json:"pass"`
	Port int    `json:"port"`
	User string `json:"user"`
}

// GetListener provides the determined listener from current configuration
func (c *Config) GetListener() string {
	return c.Listener
}

// GetRootDir will give back the root directory for all peripheral shortoftime files
func (c *Config) GetRootDir() string {
	return c.RootDir
}

// LoadFromFile will use the existing filename and load any configuration found within it
func (c *Config) LoadFromFile() (Conf, error) {
	if c.ConfigFile != "" {
		dt, err := ioutil.ReadFile(c.ConfigFile)
		if err != nil {
			return c, fmt.Errorf("Config file load error, %s", err)
		}

		json.Unmarshal(dt, &c)
	}

	return c, nil
}

// IsDebug determines whether debug mode is enabled for the current instance
func (c *Config) IsDebug() bool {
	return c.Debug
}
