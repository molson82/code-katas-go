package easyWallpaper

import "math"

var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve","thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}

func WallPaper(l, w, h float64) string {
  squareAreaNeeded := ((2*l*h) + (2*w*h)) * 1.15
  rollsNeeded := math.Floor(squareAreaNeeded / (0.52*10))

  return numbers[int(rollsNeeded) + 1]
}
