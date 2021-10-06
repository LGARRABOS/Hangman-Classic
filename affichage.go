package piscine

import "fmt"


func PrintHangmanError(Letterincorect bool , nbbeforelose int) {

	fmt.Println("Not present in the word, ", nbbeforelose ,"attempts remaining")
	
	if nbbeforelose == 9 {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("=========")

	} else if nbbeforelose == 8 {
		fmt.Println()
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")

	} else if nbbeforelose == 7 {
		fmt.Println("  +---+")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")

	} else if nbbeforelose == 6 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")

	} else if nbbeforelose == 5 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")

	} else if nbbeforelose == 4 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if nbbeforelose == 3 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
		
	} else if nbbeforelose == 2 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|\\  |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
		
	} else if nbbeforelose == 1 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" /    |")
		fmt.Println("      |")
		fmt.Println("=========")
		
	} else if nbbeforelose == 0 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  o   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" / \\  |")
		fmt.Println("      |")
		fmt.Println("=========")
		fmt.Println("The poor José is dead because of you.")
	}
}
