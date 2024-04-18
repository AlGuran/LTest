package config

import (
	"LTest/src/models"
	"github.com/mgutz/logxi/v1"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
)

var (
	Config = &models.Configuration{}
	logger = log.New("config")
)

func init() {
	configFile := "config.yml"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]) + "/../config/")

	if err != nil {
		logger.Error("Can't get startup dir", err)
	}

	f := path.Join(dir, configFile)
	logger.Info("Loading config", "file", f)

	Config, err = LoadMainConfig(&f)
	if err != nil {
		logger.Error("Error on configuration reading", err)
		os.Exit(1)
	}

	logger.Info("Config loaded")
}

func LoadMainConfig(filename *string) (*models.Configuration, error) {
	content, err := os.ReadFile(*filename)
	if err != nil {
		return nil, err
	}

	config := &models.Configuration{}
	if err := yaml.Unmarshal(content, config); err != nil {
		return nil, err
	}
	return config, nil
}
