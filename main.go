package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/agathver/coffee-beans/config"
	"github.com/agathver/coffee-beans/utils"
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
		println("Repository: ", repository.Id())

		id := repository.Id()
		r.PUT(fmt.Sprintf("/repo/%s/*artifact", id), uploadHandler(id, cfg.ContentRoot()))
		r.StaticFS(fmt.Sprintf("/repo/%s", id), gin.Dir(filepath.Join(cfg.ContentRoot(), "repositories", id), true))
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server().ListenAddress(), cfg.Server().ListenPort())
	utils.Must(r.Run(addr))
}

func uploadHandler(repository, contentRoot string) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactPath := strings.TrimLeft(c.Param("artifact"), "/")
		println(artifactPath)

		destPath := filepath.Join(contentRoot, repository, artifactPath)

		println(destPath)

		dir := filepath.Dir(destPath)

		println(dir)

		utils.Must(os.MkdirAll(dir, 0750|os.ModeDir))

		destFile, err := os.Create(destPath)
		utils.Must(err)
		defer utils.MustDo(destFile.Close)

		_, err = io.Copy(destFile, c.Request.Body)
		utils.Must(err)

		c.String(http.StatusOK, "OK")
	}
}
