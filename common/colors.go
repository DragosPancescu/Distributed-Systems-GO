package common

import "math/rand"

const (
	color_red    = "\033[0;31m"
	color_green  = "\033[0;32m"
	color_yellow = "\033[0;33m"
	color_blue   = "\033[0;34m"
	color_purple = "\033[0;35m"
	color_cyan   = "\033[0;36m"
	color_reset  = "\033[0m"
)

func Color_string(s string, c string) string {
	return c + s + color_reset
}

func Get_random_color() string {
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
