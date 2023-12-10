package main

import (
	"github.com/Xenous-Inc/finapp-api/internal/di"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/Xenous-Inc/finapp-api/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.New(config)
	server := container.GetServer()

	server.StartListening()
}
