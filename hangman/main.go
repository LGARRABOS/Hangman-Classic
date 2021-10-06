package main 

import (
	"piscine"
	"fmt"
)

func main() {
	attempts := 10
	var choice string
	c := 0
	fmt.Println("Not present in the word, ", attempts, " attempts.")
	word := piscine.RandomWord()
	tabunderscore := make([]rune, len(word))
	baseletter := piscine.LetterRandom(word)
	tabunderscore = piscine.Affichagefind(word, baseletter, tabunderscore)
	for attempts > 0 {
		choice = piscine.Choice()
		c = 0
		if piscine.AllVerif(choice) {
			if piscine.Verif_letter_in_word(word, choice) {
				tabunderscore = piscine.Affichagefind(word, choice, tabunderscore)
				for i := 0; i < len(word); i++ {
					if rune(word[i]) == tabunderscore[i] {
						c++
					}
				}
			} else {
				attempts--
				piscine.PrintHangmanError(attempts)
			}

		}
		if c == len(word) {
		 fmt.Println("Congrats !")
		 break
		}
	}
	
}