package main

import (
	"fmt"
	"github.com/recrsn/coffee-beans/repositories"

	"github.com/gin-gonic/gin"

	"github.com/recrsn/coffee-beans/config"
	"github.com/recrsn/coffee-beans/handlers"
	"github.com/recrsn/coffee-beans/utils"
)

func main() {
	cfg, err := config.Load()
	utils.Must(err)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", handlers.Ping())

	manager := repositories.NewRepositoryManager(cfg.Server.BaseURL, cfg.Repositories)

	router.GET("/", repositories.RepositoryList(manager))
	router.GET("/repositories/:id", repositories.RepositoryDetails(manager))

	for _, r := range manager.GetRepositories() {
		router.PUT(r.Path(), handlers.Upload(r.Root))
		router.StaticFS(r.StaticPath(), gin.Dir(r.Root, true))
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.ListenAddress, cfg.Server.ListenPort)
	utils.Must(router.Run(addr))
}
