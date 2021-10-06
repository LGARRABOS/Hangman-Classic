package piscine

import "fmt"

func Affichagefind(word string, letterchoose string) {
	var tabunderscore []string
	var txtletter string
	for i:= 0; i < len(word); i++ {
		tabunderscore = append(tab,"_")
	}
	for k := 0; k < len(word); k++ {
		if letterchoose == word[k] {
			tabunderscore[k] = letterchoose
			txtletter = txtletter + tabunderscore[k]
		}
		txtletter = txtletter + tabunderscore[k]
	}
	fmt.Println(txtletter)
}
