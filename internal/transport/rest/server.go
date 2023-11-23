package rest

import (
	"context"
	"fmt"
	"go-messenger/internal/config"
	"go-messenger/pkg/postgresql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	readTimeut, writeTimeout = 10 * time.Second, 10 * time.Second
)

type RestServer struct {
	httpServer *http.Server
}

func NewRestServer(ctx context.Context) RestServer {
	cfg := config.GetConfig()
	dbCfg := cfg.Db
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DbName)

	cnPool := postgresql.Init(connString, ctx)

	// place to register all storages with db

	//

	//place to register all services with storage

	//

	mainRouter := gin.Default()
	//place to register all handlers

	//

	httpServer := &http.Server{
		Addr:           ":" + cfg.AppPort,
		Handler:        mainRouter,
		ReadTimeout:    readTimeut,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	return RestServer{httpServer: httpServer}
}

func (r *RestServer) Run() error {

	go func() {
		if err := r.httpServer.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.httpServer.Shutdown(ctx)
}
