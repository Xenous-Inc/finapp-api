package main

import (
	"github.com/Xenous-Inc/finapp-api/internal/di"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/Xenous-Inc/finapp-api/internal/utils/flags"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger"
)

// @title           Finapp-api
// @version         1.0
// @description     This is a server finapp-api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   https://web.telegram.org/a/#488960669
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5051/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	logger.ConfigureZeroLogger()
	container := di.New(config)

	server := container.GetServer()

	server.StartListening()
}
