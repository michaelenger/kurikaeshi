package challenge

import (
	"fmt"
	"math/rand"

	"kurikaeshi/internal/colors"
	"kurikaeshi/internal/data"
)

// Challenge the user to write the correct rōmaji
func RunChallenge(words []data.Word) {
	for { // there is no escape
		word := words[rand.Intn(len(words))]

		var guess string
		fmt.Printf("%v: ", word.Letters)
		fmt.Scanln(&guess)

		var output string
		if guess == word.Romaji {
			output = colors.Green(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
		} else {
			output = colors.Red(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
		}

		fmt.Println(output)
	}
}
