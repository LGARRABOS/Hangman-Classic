package piscine

import "os"

func Verif_letter_in_word(word string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == os.Args[1][0] {
			return true
		}
	}
	return false
}

func Verif_letter() bool {
	if os.Args[1][0] >= 60 && os.Args[1][0] <= 122 {
		return true
	}
	return false
}

func Verif_taille_osArgs() bool {
	if len(os.Args) == 2 {
		return true
	}
	return false
}

func Verif_taille_string() bool {
	if len(os.Args[1]) == 1 {
		return true
	}
	return false
}

func lettre_utiliser() bool {
	var stock []byte
	for i := 0; i < len(stock); i++ {

	}
	stock = stock + os.Args[1][0]
}
