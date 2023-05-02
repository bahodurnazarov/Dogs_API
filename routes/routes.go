package routes

import (
	"net/http"

	"github.com/bahodurnazarov/Dogs_API/controllers"
	"github.com/bahodurnazarov/Dogs_API/cors"
	"github.com/gin-gonic/gin"
)

func Listen() {
	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	router.LoadHTMLGlob("templates/*")
	router.POST("/upload/frontPage", controllers.UploadResizeSingleFile)
	router.StaticFS("/images", http.Dir("public"))
	router.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})
	router.GET("/dog", getDogLink)
	router.POST("/addPuppy", addPuppy)
	router.GET("/getAll", getAll)
	router.DELETE("/delete", delete)
	router.Run(":8080")
}
