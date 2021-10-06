package piscine

import "fmt"

func Affichagefind(word string, letterchoose string, tabunderscore []rune) []rune {
	fmt.Println(word)
	var txtletter string
	
	for i:= 0; i < len(tabunderscore); i++ {
		if tabunderscore[i] <= 'z' && tabunderscore[i] >= 'a' {
		} else if tabunderscore[i] == '_' && word[i] == letterchoose[0] {
			tabunderscore[i] = rune(letterchoose[0])
		} else {
			tabunderscore[i] = '_'
		}
		txtletter = txtletter + string(tabunderscore[i])
	}
	fmt.Println(txtletter)
	return tabunderscore

}


