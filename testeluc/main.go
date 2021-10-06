package main

import (
	"fmt"
	"piscine"
)

func main() {
	var stock []byte
	word := "joseph"
	l := "a"
	if piscine.Verif_taille(l) {
		if piscine.Verif_letter(l) {
			if piscine.Lettre_utiliser(&stock, l) {
				if piscine.Verif_letter_in_word(l, word) {
					fmt.Println("La lettre donnée se trouve dans le mot")
				} else {
					fmt.Println("La lettre donnée ne se trouve pas dans le mot")
				}
			} else {
				fmt.Println("La lettre donnée à déja été proposer")
			}
		} else {
			fmt.Println("Votre arguments ne peux contenir qu'une lettre minuscule")
		}
	} else {
		fmt.Println("Vous ne pouvez envoyer qu'un seul argument")
	}
	l = "a"
	if piscine.Verif_taille(l) {
		if piscine.Verif_letter(l) {
			if piscine.Lettre_utiliser(&stock, l) {
				if piscine.Verif_letter_in_word(l, word) {
					fmt.Println("La lettre donnée se trouve dans le mot")
				} else {
					fmt.Println("La lettre donnée ne se trouve pas dans le mot")
				}
			} else {
				fmt.Println("La lettre donnée à déja été proposer")
			}
		} else {
			fmt.Println("Votre arguments ne peux contenir qu'une lettre minuscule")
		}
	} else {
		fmt.Println("Vous ne pouvez envoyer qu'un seul argument")
	}
}
