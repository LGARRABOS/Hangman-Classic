package main

import "piscine"

func main() {
	a := true
	b := 10
	for m := 0; m < 10;m++ {
		if a == true {
			b--
			piscine.PrintHangmanError(a,b)
		}
	}
}