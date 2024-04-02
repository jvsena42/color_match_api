package routes

import (
	"log"
	"net/http"

	"com.jvsena42/color_match/model"
	"github.com/gin-gonic/gin"
)

func getClosestColor(context *gin.Context) {

	var color model.Color

	err := context.ShouldBindJSON(&color)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		log.Println("ERROR getClosestColor: ", err)
		return
	}
	color.SetClosestColor()
	context.JSON(http.StatusOK, color)
}
