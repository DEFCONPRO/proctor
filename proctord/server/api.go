package server

import (
	"proctor/proctord/instrumentation"
	"time"

	"proctor/proctord/config"
	"proctor/proctord/logger"

	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
)

func Start() error {
	err := instrumentation.InitNewRelic()
	if err != nil {
		logger.Fatal(err)
	}
	appPort := ":" + config.AppPort()

	server := negroni.New(negroni.NewRecovery())
	router, err := NewRouter()
	if err != nil {
		return err
	}
	server.UseHandler(router)

	logger.Info("Starting server on port", appPort)

	graceful.Run(appPort, 2*time.Second, server)

	postgresClient.Close()
	logger.Info("Stopped server gracefully")
	return nil
}
