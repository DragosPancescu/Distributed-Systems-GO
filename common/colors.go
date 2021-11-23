package common

import (
	"math/rand"
	"runtime"
	"time"
)

// Global color variables
var (
	color_red    = "\033[0;31m"
	color_green  = "\033[0;32m"
	color_yellow = "\033[0;33m"
	color_blue   = "\033[0;34m"
	color_purple = "\033[0;35m"
	color_cyan   = "\033[0;36m"
	color_reset  = "\033[0m"
)

// Init function if os is windows
// because command prompt does not support colors.
func init() {
	if runtime.GOOS == "windows" {
		color_red = ""
		color_green = ""
		color_yellow = ""
		color_blue = ""
		color_purple = ""
		color_cyan = ""
		color_reset = ""
	}
}

// Function to color a string
func Color_string(input string, color string) string {
	return color + input + color_reset
}

// Select a random color to use
func Get_random_color() string {

	rand.Seed(time.Now().UnixNano())
	rng := rand.Intn(6)
	color := ""

	switch rng {
	case 0:
		color = color_red
	case 1:
		color = color_green
	case 2:
		color = color_blue
	case 3:
		color = color_yellow
	case 4:
		color = color_purple
	case 5:
		color = color_cyan
	}

	return color
}
