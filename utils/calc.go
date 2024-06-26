package utils

import (
	"math"
	"strconv"
)

func FindClosestColor(target string, colors []string) string {
	closestDistance := -1
	closestColor := ""

	for _, color := range colors {
		distance := calculateColorDistance(target, color)
		if closestDistance == -1 || distance < closestDistance {
			closestDistance = distance
			closestColor = color
		}
	}

	return closestColor
}

func calculateColorDistance(color1, color2 string) int {
	r1, g1, b1 := parseHexColor(color1)
	r2, g2, b2 := parseHexColor(color2)

	rDiff := r1 - r2
	gDiff := g1 - g2
	bDiff := b1 - b2

	// Calculate the squared differences
	squareDiff := rDiff*rDiff + gDiff*gDiff + bDiff*bDiff

	// Return the square root of the sum of squared differences (Euclidean distance)
	return int(math.Sqrt(float64(squareDiff)))
}

func parseHexColor(color string) (int, int, int) {
	r, _ := strconv.ParseInt(color[1:3], 16, 64)
	g, _ := strconv.ParseInt(color[3:5], 16, 64)
	b, _ := strconv.ParseInt(color[5:7], 16, 64)
	return int(r), int(g), int(b)
}
