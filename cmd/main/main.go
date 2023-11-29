package main

import (
	"context"
	"fmt"
	"go-messenger/internal/config"
	"go-messenger/internal/transport/rest"
	"go-messenger/pkg/postgresql"
)

func main() {
	cfg := config.GetConfig()
	dbCfg := cfg.Db
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DbName)

	cnPool := postgresql.Init(connString, context.Background())
	defer func() {
		fmt.Println("Close db pool")
		cnPool.Close()
	}()

	restServer := rest.NewRestServer(cnPool)
	err := restServer.Run()

	if err != nil {
		panic(err)
	}

}
