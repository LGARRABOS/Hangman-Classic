package piscine

import (
	"os"
	"fmt"
)
func LetterType(tabunderscore []rune) {
	if len(os.Args) == 4  && os.Args[2] == "--letterFile" {
		Ascii_art_min(tabunderscore)
	} else {
		for i := 0; i < len(tabunderscore); i++ {
			fmt.Printf(string(tabunderscore[i]))
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}