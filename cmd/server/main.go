package main

import (
	"github.com/ricardolonga/workshop-go/internal/server/http"
	"os/signal"
	"syscall"
	"os"
	"strings"
)

func main() {
	/*
	 * Storages...
	 */

	/*
	 * Services...
	 */

	/*
	 * Handler...
	 */
	handler := http.NewHandler()

	/*
	 * Server...
	 */
	server := http.New(getPort(), handler)
	server.ListenAndServe()

	/*
	 * Graceful shutdown...
	 */
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}

func getPort() string {
	if port := strings.TrimSpace(os.Getenv("PORT")); port != "" {
		return port
	}
	return "8080"
}