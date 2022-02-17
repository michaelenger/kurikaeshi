// Package data handles loading the word data.
package data

import (
	_ "embed"
	"encoding/csv"
	"regexp"
	"strings"
)

// A single word containing either hiragana or katakana morae along with the
// kanji for the word, its English translation and the phonetical rōmaji.
type Word struct {
	Morae       string
	Kanji       string
	Translation string
	Romaji      string
}

// Expression used to remove a subword
var subwordExpression *regexp.Regexp = regexp.MustCompile(`\s+\(.*?\)$`)

// Expression used to remove non-word characters
var nonWordExpression *regexp.Regexp = regexp.MustCompile(`\W`)

//go:embed embed/hiragana.csv
var hiraganaData string

//go:embed embed/katakana.csv
var katakanaData string

// Load hiragana from the embedded file.
func LoadHiragana() ([]Word, error) {
	var hiraganaWords []Word

	reader := csv.NewReader(strings.NewReader(hiraganaData))
	reader.Comma = ' '

	records, err := reader.ReadAll()
	if err != nil {
		return hiraganaWords, err
	}

	for _, record := range records {
		hiraganaWords = append(hiraganaWords, Word{
			record[0],
			record[1],
			record[2],
			record[3],
		})
	}

	return hiraganaWords, nil
}

// Load katakana from the embedded file.
func LoadKatakana() ([]Word, error) {
	var katakanaWords []Word

	reader := csv.NewReader(strings.NewReader(katakanaData))
	reader.Comma = ' '

	records, err := reader.ReadAll()
	if err != nil {
		return katakanaWords, err
	}

	for _, record := range records {
		katakanaWords = append(katakanaWords, Word{
			record[0],
			record[1],
			record[2],
			record[3],
		})
	}

	return katakanaWords, nil
}

// Sanitize a word and make it easier to match with.
func Sanitize(word string) string {
	var sanitizedWord string

	sanitizedWord = string(subwordExpression.ReplaceAll([]byte(word), []byte("")))
	sanitizedWord = strings.Replace(sanitizedWord, "ō", "ou", -1)
	sanitizedWord = strings.Replace(sanitizedWord, "ū", "uu", -1)
	sanitizedWord = string(nonWordExpression.ReplaceAll([]byte(sanitizedWord), []byte("")))

	return sanitizedWord
}
