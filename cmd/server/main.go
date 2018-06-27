package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NeowayLabs/logger"
	"github.com/ricardolonga/workshop-go/domain/user"
	"github.com/ricardolonga/workshop-go/internal/server/http"
	"github.com/ricardolonga/workshop-go/internal/storage/mongo"
)

const (
	envWorkshopServicePort = "WORKSHOP_SERVICE_PORT"
	envDatabaseName        = "DATABASE_NAME"
	envMongoURL            = "MONGO_URL"

	defaultDatabaseName = "workshop"
	defaultPort         = "8080"
)

func main() {
	mongoURL := getMongoURL()
	databaseName := getDatabaseName()

	/*
	 * Storages...
	 */
	userStorage, err := mongo.NewUserStorage(mongoURL, databaseName)
	if err != nil {
		logger.Error("error on creating a user storage instance: %q", err)
		return
	}

	/*
	 * Services...
	 */
	userService := user.NewService(userStorage)

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

func getMongoURL() string {
	return getEnvVar(envMongoURL)
}

func getDatabaseName() string {
	return getEnvVar(envDatabaseName, defaultDatabaseName)
}

func getEnvVar(envVar string, defaultValue ...string) string {
	value := os.Getenv(envVar)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}
