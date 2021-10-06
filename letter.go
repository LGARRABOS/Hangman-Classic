package piscine

import (
	"fmt"
)

func AllVerif(choice string) bool{
	var stock []byte

	if Verif_taille(choice) {
		if Verif_letter(choice) {
			if Lettre_utiliser(&stock, choice) {
				return true
			} else {
				fmt.Println("La lettre donnée à déja été proposer")
			}
		} else {
			fmt.Println("Votre arguments ne peux contenir qu'une lettre minuscule")
		}
	} else {
		fmt.Println("Vous ne pouvez envoyer qu'un seul argument")
	}
	return false
}
