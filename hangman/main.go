package main 

import (
	"piscine"
	"fmt"
)

func main() {
	attempts := 10
	lattempts := 10
	var choice string
	c := 0
	var stock []byte
	fmt.Println("Good luck, you have", attempts, " attempts.")
	word := piscine.RandomWord()
	tabunderscore := make([]rune, len(word))
	baseletter := piscine.LetterRandom(word)
	tabunderscore = piscine.Affichagefind(word, baseletter, tabunderscore)
	piscine.LetterType(tabunderscore)
	piscine.PrintHangmanError(attempts, &lattempts)
	for attempts > 0 {
		choice = piscine.Choice()
		c = 0
		if piscine.AllVerif(choice, &stock, word) {
			if piscine.Verif_letter_in_word(word, choice) {
				tabunderscore = piscine.Affichagefind(word, choice, tabunderscore)
				piscine.LetterType(tabunderscore)
				piscine.PrintHangmanError(attempts, &lattempts)
				for i := 0; i < len(word); i++ {
					if rune(word[i]) == tabunderscore[i] {
						c++
					}
				}
			} else {
				attempts--
				piscine.LetterType(tabunderscore)
				piscine.PrintHangmanError(attempts, &lattempts)
			}
			if c == len(word) {
			 fmt.Println("Congrats !")
			 break
			}
		} else if len(choice) == len(word) {
			if piscine.Complet_word(word, choice) {
				for i := 0; i < len(word); i++ {
					tabunderscore[i] = rune(word[i])
				}
				piscine.LetterType(tabunderscore)
				piscine.PrintHangmanError(attempts, &lattempts)
				fmt.Println("\n")
				fmt.Println("Congrats !")
				break
			} else {
				attempts -= 2
				piscine.LetterType(tabunderscore)
				piscine.PrintHangmanError(attempts, &lattempts)
			}
		}
	}
}