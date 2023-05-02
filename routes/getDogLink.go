package routes

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	d "github.com/bahodurnazarov/Dogs_API/db"
	"github.com/bahodurnazarov/Dogs_API/models"
	lg "github.com/bahodurnazarov/Dogs_API/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // add this
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
	db := d.ConnDB()
	if data.AppVersion != "" {
		_, err := db.Exec("INSERT into dogs VALUES ($1, $2, $3)", data.Name, data.Age, data.AppVersion)
		if err != nil {
			lg.Errl.Fatalf("An error occured while executing query: %v", err)
		}
	}

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		lg.Errl.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
