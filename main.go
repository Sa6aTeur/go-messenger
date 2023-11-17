package main

import (
	"go-messenger/config"
	"go-messenger/router"
)

func main() {
	cfg := config.GetConfig()
	r := router.SetupRouter()

	r.Run(":" + cfg.AppPort)

}
