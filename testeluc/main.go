package main

import (
	"fmt"
	"piscine"
)

func main() {
	word := "joseph"
	if piscine.Verif_taille_osArgs() {
		if piscine.Verif_taille_string() {
			if piscine.Verif_letter() {
				if piscine.lettre_utiliser() {
					if piscine.Verif_letter_in_word(word) {
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
			fmt.Println("Votre arguments ne doit contenir qu'un seul carractére")
		}
	} else {
		fmt.Println("Vous ne pouvez envoyer qu'un seul argument")
	}
}
