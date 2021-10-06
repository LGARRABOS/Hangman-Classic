package piscine

<<<<<<< HEAD
func Verif_letter_in_word(word string, letter string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == letter[0] {
=======
func Verif_letter_in_word(word string, choice string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == choice[0] {
>>>>>>> 32df850e2b2e74a1ce1a5342c75624de2f48881a
			return true
		}
	}
	return false
}

<<<<<<< HEAD
func Verif_letter(letter string) bool {
	if letter[0] >= 60 && letter[0] <= 122 {
=======
func Verif_letter(choice string) bool {
	if choice[0] >= 60 && choice[0] <= 122 {
>>>>>>> 32df850e2b2e74a1ce1a5342c75624de2f48881a
		return true
	}
	return false
}

<<<<<<< HEAD
func Verif_taille(letter string) bool {
	if len(letter) == 1 {
=======
func Verif_taille(choice string) bool {
	if len(choice) == 1 {
>>>>>>> 32df850e2b2e74a1ce1a5342c75624de2f48881a
		return true
	}
	return false
}

<<<<<<< HEAD
func Lettre_utiliser(use *[]byte, letter string) bool {
	a := *use
	for i := 0; i < len(*use); i++ {
		if a[i] == letter[0] {
			return false
		}
	}
	*use = append(*use, letter[0])
=======
func Lettre_utiliser(use *[]byte, choice string) bool {
	a := *use
	for i := 0; i < len(*use); i++ {
		if a[i] == choice[0] {
			return false
		}
	}
	*use = append(*use, choice[0])
>>>>>>> 32df850e2b2e74a1ce1a5342c75624de2f48881a
	return true
}
