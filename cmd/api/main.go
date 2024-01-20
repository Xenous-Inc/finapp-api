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
	sessionId, err := clientOrg.Login(&orgfaclient.LoginInput{
		Login:    "226292",
		Password: "Oeia7299",
	})

	if err != nil {
		fmt.Println(err)
	} else {
		_ = sessionId
		fmt.Println("jopa:", sessionId)
	}
	//Logged in successfully with id
	// cook, err := clientOrg.GetMyGroup(&orgfaclient.AuthSession{
	// 	SessionId: id,
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	_ = cook
	// 	fmt.Println("Logged in successfully with cook:", cook)

	// }

	// zachetka, err := clientOrg.GetRecordBook(&orgfaclient.AuthSession{
	// 	SessionId: sessionId,
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Logged in successfully with zachetka:", zachetka)
	// }

	// miniProfile, err := clientOrg.GetMiniProfile(&orgfaclient.GetMiniProfileInput{
	// 	AuthSession: &orgfaclient.AuthSession{
	// 		SessionId: id,
	// 	},
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Logged in successfully with zachetka:", miniProfile)
	// }

	// profile, err := clientOrg.GetProfile(&orgfaclient.GetProfileInput{
	// 	AuthSession: &orgfaclient.AuthSession{
	// 		SessionId: id,
	// 	},
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Logged in successfully with zachetka:", profile)
	// }

	// order, err := clientOrg.GetOrder(&orgfaclient.GetOrderInput{
	// 	AuthSession: &orgfaclient.AuthSession{
	// 		SessionId: sessionId,
	// 	},
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Logged in successfully with order:", order)
	// }

	// DONT WORK
	// studentCard, err := clientOrg.GetStudentCard(&orgfaclient.GetStudentCardInput{
	// 	AuthSession: &orgfaclient.AuthSession{
	// 		SessionId: sessionId,
	// 	},
	// 	ProfileId: "93491",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Logged in successfully with studentCard:", studentCard)
	// }

	// studentPlan, err := clientOrg.GetStudyPlan(&orgfaclient.GetStudyPlanInput{
	// 	AuthSession: &orgfaclient.AuthSession{
	// 		SessionId: sessionId,
	// 	},
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Logged in successfully with student plan:", studentPlan)
	// }

	server.StartListening()
}
