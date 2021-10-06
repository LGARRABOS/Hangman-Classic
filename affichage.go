package piscine

import (
	"fmt"
)

func PrintHangmanError(error bool , lose int) {
	if lose == 9 {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("=========")
	} else if lose == 8 {
		fmt.Println()
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if lose == 7 {
		fmt.Println("  +---+")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if lose == 6 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if lose == 5 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")

	} else if lose == 4 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if lose == 3 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
		
	} else if lose == 2 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|\\  |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
		
	} else if lose == 1 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" /    |")
		fmt.Println("      |")
		fmt.Println("=========")
		
	} else if lose == 0 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" / \\  |")
		fmt.Println("      |")
		fmt.Println("=========")
	}
}
