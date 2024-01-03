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
	schedule, err := client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   "111144",
		StartDate: "2023.12.11",
		EndDate:   "2023.12.11",
	})
	teacher, err := client.GetTeacher(&ruzfaclient.GetTeacherInput{
		TeacherTerm: "Бердышев",
	})
	scheduleTeacher, err := client.GetTeacherSchedule(&ruzfaclient.GetTeacherScheduleInput{
		Id:        "00000000-0001-2345-6789-000000102845",
		StartDate: "2023.12.11",
		EndDate:   "2023.12.17",
	})
	auditorium, err := client.GetAuditorium(&ruzfaclient.GetAuditoriumInput{
		AuditoriumTerm: "304",
	})
	scheduleAuditorium, err := client.GetAuditoriumSchedule(&ruzfaclient.GetAuditoriumScheduleInput{
		Id:        "3127",
		StartDate: "2023.12.18",
		EndDate:   "2023.12.24",
	})

	if err != nil {
		log.Fatal(err)
	}
	_ = groups
	_ = schedule
	_ = teacher
	_ = scheduleTeacher
	_ = scheduleAuditorium
	_ = auditorium

	fmt.Println(scheduleAuditorium[1].String())
	server.StartListening()
}
