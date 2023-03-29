package routes

import (
	"net/http"

	"github.com/bahodurnazarov/Dogs_API/cors"
	"github.com/gin-gonic/gin"
)

func Listen() {
	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	router.LoadHTMLGlob("templates/*")
	router.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})
	router.GET("/dog", getDogLink)

	router.Run(":8080")
}
