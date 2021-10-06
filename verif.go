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

func Lettre_utiliser(use *[]byte) bool {
	a := *use
	for i := 0; i < len(*use); i++ {
		if a[i] == os.Args[1][0] {
			return false
		}
	}
	*use = append(*use, os.Args[1][0])
	return true
}
