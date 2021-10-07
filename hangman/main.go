package main 

import (
	"piscine"
	"fmt"
	"os"
)

func main() {
	attempts := 10
	var choice string
	c := 0
	var stock []byte
	fmt.Println("Good luck, you have", attempts, " attempts.")
	if len(os.Args) == 2 {
		word := piscine.RandomWord()
		tabunderscore := make([]rune, len(word))
		baseletter := piscine.LetterRandom(word)
		tabunderscore = piscine.Affichagefind(word, baseletter, tabunderscore)
		piscine.Ascii_art_min(tabunderscore)
		for attempts > 0 {
			choice = piscine.Choice()
			c = 0
			if piscine.AllVerif(choice, &stock, word) {
				if piscine.Verif_letter_in_word(word, choice) {
					tabunderscore = piscine.Affichagefind(word, choice, tabunderscore)
					piscine.Ascii_art_min(tabunderscore)
					for i := 0; i < len(word); i++ {
						if rune(word[i]) == tabunderscore[i] {
							c++
						}
					}
				} else {
					attempts--
					piscine.PrintHangmanError(attempts)
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
					piscine.Ascii_art_min(tabunderscore)
					fmt.Println("\n")
					fmt.Println("Congrats !")
					break
				} else {
					attempts -= 2
					piscine.PrintHangmanError(attempts)
				}
			}
		}
	} else {
		fmt.Println("pas le bon nombre de fichier")
	}
}