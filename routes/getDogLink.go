package routes

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/bahodurnazarov/Dogs_API/models"
	lg "github.com/bahodurnazarov/Dogs_API/utils"
	"github.com/gin-gonic/gin"
)

func getDogLink(c *gin.Context) {

	lg.Server.SetFlags(log.Lshortfile)

	appVersion := models.GetDogLink()
	dogAge := rand.Intn(13)
	dogName := models.PetName()
	lg.Server.Println(dogAge)

	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		lg.Errl.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	data := struct {
		AppVersion string
		Name       string
		Age        int
	}{
		AppVersion: appVersion,
		Name:       dogName,
		Age:        dogAge,
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		lg.Errl.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
