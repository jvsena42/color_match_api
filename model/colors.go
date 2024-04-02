package model

import (
	"com.jvsena42/color_match/utils"
)

type Color struct {
	ColorHex     string `json:"color_hex"`
	ClosestColor string `json:"closest_color"`
}

func (color *Color) GetClosestColor() {
	colorList := []string{"#FF4444", "#00FF00", "#FF0044", "#4444FF"}

	findedColor := utils.FindClosestColor(color.ColorHex, colorList)

	color.ClosestColor = findedColor
}
