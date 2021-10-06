package piscine 

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)
func RandomWord() string {
	rand.Seed(time.Now().UnixNano())
	var wordtab []string
	var word string
	f, err := os.Open("word.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordtab = append(wordtab, scanner.Text())
    }
	word = wordtab[rand.Intn(len(wordtab))]
	return word
	
}

func LetterRandom(word) string {
	var letter string
	letter = word[rand.Intn(len(word))]
}
