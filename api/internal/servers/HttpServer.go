package servers

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type HttpServer struct {
	port    string
	context context.Context
}

func NewHttpServer(ctx context.Context) (*HttpServer, error) {

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", ":3000")
	}

	port := os.Getenv("PORT")

	return &HttpServer{
		port:    port,
		context: ctx,
	}, nil
}

func (server *HttpServer) Run() error {
	fmt.Println("HttpServer is running on port: ", server.port)

	err := http.ListenAndServe(server.port, nil)

	if err != nil {
		return err
	}

	return nil
}
