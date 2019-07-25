package main

import (
	"github.com/ricardolonga/workshop-go/domain/user"
	"os"
	"os/signal"
	"syscall"

	"github.com/ricardolonga/workshop-go/internal/server/http"
)

const (
	envWorkshopServicePort = "WORKSHOP_SERVICE_PORT"
	defaultPort            = "8080"
)

func main() {
	userService := user.NewService()

	/*
	 * Handler...
	 */
	handler := http.NewHandler(userService)

	/*
	 * Server...
	 */
	server := http.New(getApplicationPort(), handler)
	server.ListenAndServe()

	/*
	 * Graceful shutdown...
	 */
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}

func getApplicationPort() string {
	return getEnvVar(envWorkshopServicePort, defaultPort)
}

func getEnvVar(envVar string, defaultValue ...string) string {
	value := os.Getenv(envVar)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}
