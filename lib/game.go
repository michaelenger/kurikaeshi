package lib

import (
	"fmt"
	"math/rand"
)

// Challenge the user to write the correct rōmaji
func Challenge(words []Word) {
	for { // there is no escape
		word := words[rand.Intn(len(words))]

		var guess string
		fmt.Printf("%v: ", word.Letters)
		fmt.Scanln(&guess)

		if guess == word.Romaji {
			fmt.Printf("CORRECT! %v (%v): %v\n", word.Letters, word.Kanji, word.Romaji)
		} else {
			fmt.Printf("WRONG! %v (%v): %v\n", word.Letters, word.Kanji, word.Romaji)
		}
	}
}
