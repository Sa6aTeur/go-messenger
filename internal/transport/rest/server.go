package rest

import (
	"context"
	"fmt"
	"go-messenger/internal/config"
	"go-messenger/internal/service"
	"go-messenger/internal/storage/psql"
	"go-messenger/internal/transport/rest/handler"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	readTimeut, writeTimeout = 10 * time.Second, 10 * time.Second
)

type RestServer struct {
	httpServer *http.Server
}

func NewRestServer(dbPool *pgxpool.Pool) RestServer {
	cfg := config.GetConfig()

	storages := psql.New(dbPool)

	//place to register all services with storages
	userService := service.NewUserService(storages.UserStorage)

	//place to register all handlers with services
	userH := handler.NewUserHandler(userService)

	mainRouter := gin.Default()
	
	//place to register all endpoints
	userH.RegisterEndpoints(mainRouter)

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
