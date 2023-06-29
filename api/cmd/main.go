package main

import (
	"context"

	"github.com/gt2rz/status-order-update/api/internal/servers"
)

func main() {
	httpServer, err := servers.NewHttpServer(context.Background())
	if err != nil {
		// throw error
		panic(err)
	}

	httpServer.Run()
}
