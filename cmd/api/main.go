package main

import (
	"log"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/di"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/Xenous-Inc/finapp-api/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.New(config)
	server := container.GetServer()

	client := ruzfaclient.NewClient(&http.Client{}, "https://ruz.fa.ru/api/")
	groups, err := client.GetGroups(&ruzfaclient.GetGroupsInput{
		GroupTerm: "ФФ22",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(groups)

	server.StartListening()
}
