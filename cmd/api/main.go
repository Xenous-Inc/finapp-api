package main

import (
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
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
	clientOrg := orgfaclient.NewClient(&http.Client{}, "https://org.fa.ru/")
	server := container.GetServer(client, clientOrg)
	id, err := clientOrg.Login(&orgfaclient.LoginInput{
		Login:    "226292",
		Password: "Oeia7299",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Logged in successfully with id:", id)
	}

	cook, err := clientOrg.GetMyGroup(&orgfaclient.AuthSession{
		SessionId: id,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Logged in successfully with cook:", cook)
	}
	
	server.StartListening()
}
