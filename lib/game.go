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

		var output string
		if guess == word.Romaji {
			output = Green(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
		} else {
			output = Red(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
		}

		fmt.Println(output)
	}
}
