package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)

	})


	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html",nil)

	})
	r.Static("/public","./public")

	r.Run(":4200")
}
