package routes

import (
	d "github.com/bahodurnazarov/Dogs_API/db"
	"github.com/gin-gonic/gin"
)

func delete(c *gin.Context) {
	thsImg := c.Query("dogimage")
	db := d.ConnDB()

	db.Exec("DELETE from dogs WHERE dogimage=$1", thsImg)
}
