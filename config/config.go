package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	configFile string
	Options    *Options
	logger     *zap.Logger
}

func New(configFile string, logger *zap.Logger) (*Config, error) {
	c := &Config{
		configFile: configFile,
	}

	yamFile, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error("Unable to read configuration file", zap.Error(err))
		return nil, err
	}

	err = yaml.Unmarshal(yamFile, &c.Options)
	if err != nil {
		logger.Error("Unable to read configuration file", zap.Error(err))
		return nil, err
	}

	return c, nil
}
