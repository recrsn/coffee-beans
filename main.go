package main

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/agathver/coffee-beans/app/config"
	"github.com/agathver/coffee-beans/app/handlers"
	"github.com/agathver/coffee-beans/app/utils"
)

func main() {
	cfg, err := config.Load()
	utils.Must(err)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	for _, repository := range cfg.Repositories() {
		id := repository.Id()
		r.PUT(fmt.Sprintf("/repo/%s/*artifact", id), handlers.Upload(id, cfg.ContentRoot()))
		r.StaticFS(fmt.Sprintf("/repo/%s", id), gin.Dir(filepath.Join(cfg.ContentRoot(), "repositories", id), true))
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server().ListenAddress(), cfg.Server().ListenPort())
	utils.Must(r.Run(addr))
}
