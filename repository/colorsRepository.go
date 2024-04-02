package repository

import (
	"com.jvsena42/color_match/utils"
)

func GetClosestColor(color string) string {
	colorList := []string{"#FF4444", "#00FF00", "#FF0044", "#4444FF"}

	foundColor := utils.FindClosestColor(color, colorList)

	return foundColor
}
