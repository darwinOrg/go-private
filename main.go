package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var (
	privateDomain        = os.Getenv("PRIVATE_DOMAIN")
	privateLibPathPrefix = os.Getenv("PRIVATE_LIB_PATH_PREFIX")
	privateGitPathPrefix = os.Getenv("PRIVATE_GIT_PATH_PREFIX")
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("resources/*")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "ok")
	})
	r.GET(privateLibPathPrefix+"/:mod", func(c *gin.Context) {
		mod := c.Param("mod")
		c.HTML(http.StatusOK, "default.tmpl", gin.H{
			"privateDomain":        privateDomain,
			"privateLibPathPrefix": privateLibPathPrefix,
			"privateGitPathPrefix": privateGitPathPrefix,
			"mod":                  mod,
		})
	})
	_ = r.Run(":8080")
}
