package config

import (
	"errors"

	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
	"github.com/spf13/viper"
)

const (
	ENV_MODE_DEVELOPMENT = iota + 1
	ENV_MODE_PRODUCTION  = iota + 1
	ENV_MODE_STAGE       = iota + 1
)

const (
	envModeDevelopmentStr = "development"
	envModeProductionStr  = "production"
	envModeStageStr       = "stage"
)

type Config struct {
	EnvMode   uint8
	Host      string
	Port      uint16
	JwtSecret string
}

func LoadConfig(envMode, path string) (*Config, error) {
	mode, err := validateEnvMode(envMode)
	if err != nil {
		log.Error(err, "Internal", "utils/config LoadConfig")
		return nil, err
	}

	config := new(Config)

	// Getting configuration file
	viper.SetConfigFile(path)

	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err, "Internal", "utils/config LoadConfig")
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Error(err, "Internal", "utils/config LoadConfig")
		return nil, err
	}

	config.EnvMode = mode

	return config, nil
}

func MustLoadConfig(envMode, path string) *Config {
	config, err := LoadConfig(envMode, path)
	if err != nil {
		log.Error(err, "Internal", "utils/config MustLoadConfig")
		panic(err)
	}

	return config
}

func validateEnvMode(envMode string) (uint8, error) {
	var mode uint8
	switch envMode {
	case envModeDevelopmentStr:
		mode = ENV_MODE_DEVELOPMENT
	case envModeProductionStr:
		mode = ENV_MODE_PRODUCTION
	case envModeStageStr:
		mode = ENV_MODE_STAGE
	default:
		log.Warn("Internal", "utils/config validateEnvMode")
		return mode, errors.New("Unknown environment mode")
	}

	return mode, nil
}
