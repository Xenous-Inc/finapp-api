package flags

import (
	"errors"
	"flag"

	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
)

const (
	configPathFlag = "config-path"
	envModeFlag    = "env-mode"
)

type CMDFlags struct {
	ConfigPath string
	EnvMode    string
}

func ParseFlags() (*CMDFlags, error) {
	configPath := flag.String(configPathFlag, "", "Configuration file path")
	envMode := flag.String(envModeFlag, "", "Environment mode")
	flag.Parse()

	if *configPath == "" {
		log.Warn("Internal", "flags ParseFlags")
		return nil, errors.New("Configuration file path was not found in application flags")
	}

	if *envMode == "" {
		log.Warn("Internal", "flags ParseFlags")
		return nil, errors.New("Environment mode was not found in application flags")
	}

	return &CMDFlags{ConfigPath: *configPath, EnvMode: *envMode}, nil
}

func MustParseFlags() *CMDFlags {
	flags, err := ParseFlags()
	if err != nil {
		log.Error(err, "Internal", "flags MustParseFlags")
		panic(err)
	}

	return flags
}
