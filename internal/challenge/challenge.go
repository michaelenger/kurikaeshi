// Package challenge handles the word challenge.
package challenge

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"kurikaeshi/internal/colors"
	"kurikaeshi/internal/data"
)

var subwordExpression *regexp.Regexp = regexp.MustCompile(`\s+\(.*?\)$`)

// Challenge the user to write the correct rōmaji
func RunChallenge(words []data.Word) {
	var guess string
	var output string

	for { // there is no escape
		word := words[rand.Intn(len(words))]

		fmt.Printf("%v: ", word.Letters)
		fmt.Scanln(&guess)

		if guess == sanitize(word.Romaji) {
			output = colors.Green(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
		} else {
			output = colors.Red(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
		}

		fmt.Println(output)
	}
}

// Sanitize a word and make it easier to match with.
func sanitize(word string) string {
	var sanitizedWord string

	sanitizedWord = string(subwordExpression.ReplaceAll([]byte(word), []byte("")))
	sanitizedWord = strings.Replace(sanitizedWord, "-", "", -1)
	sanitizedWord = strings.Replace(sanitizedWord, "ō", "ou", -1)
	sanitizedWord = strings.Replace(sanitizedWord, "ū", "uu", -1)

	return sanitizedWord
}
