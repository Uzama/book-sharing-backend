package config

import (
	"fmt"
	"os"
	"strconv"
)

func Parse() (Config, error) {
	
	app := &App{}
	db := &Database{}

	err := app.Parse()
	if err != nil {
		return Config{}, err
	}

	err = db.Parse()
	if err != nil {
		return Config{}, err
	}

	configs := Config{
		App:      *app,
		Database: *db,
	}

	return configs, nil
}

func getEnvInt(key string) (int, error) {
	s := os.Getenv(key)
	if s == "" {
		return 0, fmt.Errorf("%s is empty", key)
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func getEnvString(key string) (string, error) {
	s := os.Getenv(key)
	if s == "" {
		return "", fmt.Errorf("%s is empty", key)
	}
	return s, nil
}
