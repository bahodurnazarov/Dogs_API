package controllers

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	lg "github.com/bahodurnazarov/Dogs_API/utils"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func UploadResizeSingleFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		lg.Errl.Println("file err : %s")
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://localhost:8080/images/frontPage/" + filename

	imageFile, _, err := image.Decode(file)
	if err != nil {
		lg.Errl.Println(err)
	}
	src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
	err = imaging.Save(src, fmt.Sprintf("public/frontPage/%v", filename))
	if err != nil {
		lg.Errl.Println("failed to save image:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
	log.Println(filePath)

}
