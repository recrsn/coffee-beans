package repositories

import (
	"github.com/gin-gonic/gin"
)

func RepositoryList(manager *RepositoryManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(200, "repositories.html", gin.H{
			"repositories": manager.GetRepositories(),
		})
	}
}

func RepositoryDetails(manager *RepositoryManager) func(c *gin.Context) {
	return func(c *gin.Context) {
		repo, found := manager.GetRepository(c.Param("id"))

		if !found {
			c.AbortWithStatus(404)
			return
		}

		c.HTML(200, "repository.html", repo)
	}
}
