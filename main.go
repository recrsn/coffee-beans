package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const root = "data/repository"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.PUT("/repository/*artifact", func(c *gin.Context) {
		artifactPath := strings.TrimLeft(c.Param("artifact"), "/")
		println(artifactPath)

		destPath := filepath.Join(root, artifactPath)

		println(destPath)

		dir := filepath.Dir(destPath)

		println(dir)

		must(os.MkdirAll(dir, 0750|os.ModeDir))

		destFile, err := os.Create(destPath)
		must(err)
		defer destFile.Close()

		_, err = io.Copy(destFile, c.Request.Body)
		must(err)

		c.String(http.StatusOK, "OK")
	})

	r.StaticFS("/repository", gin.Dir("data/repository", true))
	err := r.Run()

	if err != nil {
		panic(err)
	}
}
