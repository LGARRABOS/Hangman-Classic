package piscine

import "fmt"

func Affichagefind(word string, letterchoose string) {
	var tabunderscore []string
	var txtletter string
	var tabl []rune
	for i := 0; i < len(letterchoose); i++ {
		tabl = append(tabl, rune(letterchoose[0]))
	}
	for i:= 0; i < len(word); i++ {
		tabunderscore = append(tabunderscore,"_")
	}
	for k := 0; k < len(word); k++ {
		if tabl[0] == rune(word[k]) {
			tabunderscore[k] = letterchoose
			txtletter = txtletter + tabunderscore[k]
		}
		txtletter = txtletter + tabunderscore[k]
	}
	fmt.Println(txtletter)
}
