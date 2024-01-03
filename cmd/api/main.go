package main

import (
	"fmt"
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
	schedule, err := client.GetSchedule(&ruzfaclient.GetScheduleInput{
		GroupId: "111144",
		StartDate: "2023.12.11",
		EndDate: "2023.12.11",
	})
	
	if err != nil {
		log.Fatal(err)
	}
	_ = groups

	fmt.Println(schedule[0].String())
	//log.Println(schedule[0])

	server.StartListening()
}
