package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port int `json:"port"`
}

func SetEnvironment() (Config, error) {
	var config = Config{
		Port: 8080,
	}
	var err error
	if os.Getenv("port") != "" {
		config.Port, err = strconv.Atoi(os.Getenv("port"))
		if err != nil {
			return config, err
		}
	}

	return config, nil
}