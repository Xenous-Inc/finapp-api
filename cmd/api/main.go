package main

import (
	"github.com/Xenous-Inc/finapp-api/internal/di"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/Xenous-Inc/finapp-api/internal/utils/flags"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger"
)

// @title           FinAppAPI
// @version         1.0
// @description     This is a FinAppAPI Swaggers documentation. For any suggestions, questions and deals write me in Telegram: @dr3dnought
// @termsOfService  http://swagger.io/terms/

// @contact.name   Alexander Drednought
// @contact.email  ask@dr3dnought.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5051/

// @securityDefinitions.basic  BasicAuth
func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	logger.ConfigureZeroLogger()
	container := di.New(config)

	server := container.GetServer()
	server.StartListening()
}
