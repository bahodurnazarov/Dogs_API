package routes

import (
	"fmt"
	"log"

	d "github.com/bahodurnazarov/Dogs_API/db"
	lg "github.com/bahodurnazarov/Dogs_API/utils"
	"github.com/gin-gonic/gin"
)

type dog struct {
	Name  string
	Age   string
	Image string
}

func addPuppy(c *gin.Context) {
	var data dog
	lg.Server.SetFlags(log.Lshortfile)
	data.Name = c.PostForm("name")
	data.Age = c.PostForm("age")
	data.Image = c.PostForm("image")
	fmt.Printf("name: %s; age: %s; image: %s  ", data.Name, data.Age, data.Image)
	db := d.ConnDB()
	//file := GetFile()
	if data.Name != "" && data.Age != "" {
		_, err := db.Query("INSERT into dogs VALUES ($1, $2, $3)", data.Name, data.Age, data.Image)
		if err != nil {
			lg.Errl.Fatalf("An error occured while executing query: %v", err)
		}
	}
	getAll(c)
}
