package model

import (
	"com.jvsena42/color_match/repository"
)

type Color struct {
	ColorHex     string `json:"color_hex"`
	ClosestColor string `json:"closest_color"`
}

func (color *Color) SetClosestColor() {
	color.ClosestColor = repository.GetClosestColor(color.ColorHex)
}
