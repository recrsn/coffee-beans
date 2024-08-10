package main

import (
	"fmt"
	"github.com/recrsn/coffee-beans/repositories"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/recrsn/coffee-beans/config"
	"github.com/recrsn/coffee-beans/handlers"
	"github.com/recrsn/coffee-beans/utils"
)

var (
	version = "dev"
	commit  = "HEAD"
)

func main() {
	log.Printf("Coffee Beans %s (%s-%s)\n", version, commit, buildMode)

	gin.SetMode(buildMode)

	cfg, err := config.Load()
	utils.Must(err)

	router := gin.Default()

	// version is a build-time variable, so it's always set
	if buildMode == "release" {
		router.SetHTMLTemplate(template.Must(template.New("").ParseFS(embedFS, "templates/*.html")))

		static, err := fs.Sub(embedFS, "static")
		utils.Must(err)

		router.StaticFS("/static", http.FS(static))
	} else {
		log.Printf("Running in development mode, loading templates and assets from filesystem\n")
		router.LoadHTMLGlob("templates/*.html")
		router.Static("/static", "static")
	}

	router.GET("/health", handlers.HealthCheck())

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
