package main

import (
	"piscine"
)

func main() {
	var g string
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := 0; i < len(alphabet); i++ {
		g = alphabet[i]
		piscine.Ascii_art_min(g)
	}
}
