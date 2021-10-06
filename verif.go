package piscine

func Verif_letter_in_word(word string, letter string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == letter[0] {
			return true
		}
	}
	return false
}

func Verif_letter(letter string) bool {
	if letter[0] >= 60 && letter[0] <= 122 {
		return true
	}
	return false
}

func Verif_taille(letter string) bool {
	if len(letter) == 1 {
		return true
	}
	return false
}

func Lettre_utiliser(use *[]byte, letter string) bool {
	a := *use
	for i := 0; i < len(*use); i++ {
		if a[i] == letter[0] {
			return false
		}
	}
	*use = append(*use, letter[0])
	return true
}
