package handlers

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/agathver/coffee-beans/server/utils"
	"github.com/gin-gonic/gin"
)

// Upload creates a new handler to handle artifact uploads
func Upload(contentRoot string) func(c *gin.Context) {
	return func(c *gin.Context) {
		artifactPath := strings.TrimLeft(c.Param("artifact"), "/")

		if isPathInvalid(artifactPath) {
			c.String(http.StatusBadRequest, "Invalid artifact path\n")
			return
		}

		destPath := filepath.Clean(filepath.Join(contentRoot, artifactPath))

		dir := filepath.Dir(destPath)

		utils.Must(os.MkdirAll(dir, 0750|os.ModeDir))

		destFile, err := os.Create(destPath)
		utils.Must(err)
		defer utils.MustDo(destFile.Close)

		_, err = io.Copy(destFile, c.Request.Body)
		utils.Must(err)

		c.String(http.StatusCreated, "OK\n")
	}
}

func isPathInvalid(artifactPath string) bool {
	match, _ := regexp.MatchString(`[\w][\w/.\-][\w]`, artifactPath)

	if !match {
		return true
	}

	if artifactPath != path.Clean(artifactPath) {
		return true
	}

	return false
}
