package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ricardolonga/workshop-go/domain/user"
	"github.com/ricardolonga/workshop-go/internal/server/http"
)

//func main() {
//	user := domain.User{Name: "Ricardo", Age: 31, Phones: []string{"48999792240"}}
//	fmt.Printf("User: %+v", user)
//}

func main() {
	/*
	 * Storages...
	 */

	/*
	 * Services...
	 */
	userService := user.NewService( /*referência de um userStorage*/ )

	/*
	 * Handler...
	 */
	handler := http.NewHandler(userService)

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
