package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	word := "joseph"
	if os.Args[1][0] >= 60 && os.Args[1][0] <= 122 {
		if piscine.Verif(word) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	} else {
		fmt.Println("Le carractére rentrée n'est pas accepter")
	}
}
