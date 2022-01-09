// Package data handles loading data from CSV files.
package data

import (
	_ "embed"
	"strings"
	"encoding/csv"
)

// A single word containing either hiragana or katakana letters along with the
// kanji for the word, its English translation and the phonetical rōmaji.
type Word struct {
	Letters     string
	Kanji       string
	Translation string
	Romaji      string
}

//go:embed embed/hiragana.csv
var hiraganaData string

// Load hiragana from the embedded file.
func LoadHiragana() ([]Word, error) {
	var hiraganaWords []Word

	reader := csv.NewReader(strings.NewReader(hiraganaData))
	reader.Comma = ' '

    records, err := reader.ReadAll()
    if err != nil {
    	return hiraganaWords, err
    }

    for _, record := range(records) {
    	hiraganaWords = append(hiraganaWords, Word{
    		record[0],
    		record[1],
    		record[2],
    		record[3],
    	})
    }

	return hiraganaWords, nil
}
