package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("resources/*")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "ok")
	})
	r.GET("/libs/:mod", func(c *gin.Context) {
		mod := c.Param("mod")
		c.HTML(http.StatusOK, "default.tmpl", gin.H{"mod": mod})
	})

	r.RunTLS(":443", "./certs/server.crt", "./certs/server.key")
	//r.Run(":8080")
}
