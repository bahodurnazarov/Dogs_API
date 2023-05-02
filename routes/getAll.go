package routes

import (
	"html/template"
	"log"
	"net/http"

	d "github.com/bahodurnazarov/Dogs_API/db"
	lg "github.com/bahodurnazarov/Dogs_API/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // add this
)

var (
	DogAge   int
	DogName  string
	DogImage string
)

type data1 struct {
	Name  []string
	Age   []int
	Image []string
}

func getAll(c *gin.Context) {
	var Dnm []string
	var Dge []int
	var Dimg []string
	lg.Server.SetFlags(log.Lshortfile)
	var data data1
	tmpl, err := template.ParseFiles("templates/main1.html")
	if err != nil {
		lg.Errl.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	db := d.ConnDB()
	rows, err := db.Query("SELECT dogname, dogage, dogimage FROM dogs")
	if err != nil {
		lg.Errl.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&DogName, &DogAge, &DogImage)
		if err != nil {
			lg.Errl.Fatal(err)
		}
		Dnm = append(Dnm, DogName)
		Dge = append(Dge, DogAge)
		Dimg = append(Dimg, DogImage)

		data = data1{
			Name:  Dnm,
			Age:   Dge,
			Image: Dimg,
		}

	}
	
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		lg.Errl.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
