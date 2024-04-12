package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Address       string `json:"address"`
	TimeoutDB     int    `json:"timeoutDB"`
	Env           string `json:"env"`
	DriverName    string `json:"driverName"`
	ConnStr       string `json:"connStr"`
	MigrationsDir string `json:"migrationsDir"`
}

func MustLoad(configPath string) *Config {
	configFile, err := os.Open(configPath)

	if err != nil {
		panic(err.Error() + ": " + configPath)
	}

	var cfg Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&cfg); err != nil {
		panic(err.Error())
	}

	return &cfg
}
