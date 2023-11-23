package main

import (
	"context"
	"go-messenger/internal/transport/rest"
)

func main() {

	restServer := rest.NewRestServer(context.Background())
	err := restServer.Run()
	if err != nil {
		panic(err)
	}

}
