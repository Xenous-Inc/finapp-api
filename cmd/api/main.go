package main

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/di"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/Xenous-Inc/finapp-api/internal/utils/flags"
)

// @title Finapp-api

// @host localhost:5555
// @BasePath /

func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.New(config)

	client := ruzfaclient.NewClient(&http.Client{}, "https://ruz.fa.ru/api/")
	server := container.GetServer(client)
	server.StartListening()
}
