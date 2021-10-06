package main 

import (
	"piscine"
	"fmt"
)

func main() {
	attempts := 10
	var choice string
	fmt.Println("Not present in the word, ", attempts, " attempts.")
	word := piscine.RandomWord()
	baseletter := piscine.LetterRandom(word)
	piscine.Affichagefind(word, baseletter)
	for attempts > 0 {
		choice = piscine.Choice()
		if piscine.AllVerif(choice) {
			if piscine.Verif_letter_in_word(word, choice) {
				piscine.Affichagefind(word, choice)
			} else {
				attempts--
				piscine.PrintHangmanError(attempts)
			}

		}
	}
	
}