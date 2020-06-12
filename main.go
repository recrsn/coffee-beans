package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/agathver/coffee-beans/app/config"
	"github.com/agathver/coffee-beans/app/handlers"
	"github.com/agathver/coffee-beans/app/utils"
)

func main() {
	cfg, err := config.Load()
	utils.Must(err)

	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	for _, repository := range cfg.Repositories {
		r.PUT(fmt.Sprintf("/repo/%s/*artifact", repository.Id), handlers.Upload(repository.Root))
		r.StaticFS(fmt.Sprintf("/repo/%s", repository.Id), gin.Dir(repository.Root, true))
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.ListenAddress, cfg.Server.ListenPort)
	utils.Must(r.Run(addr))
}
