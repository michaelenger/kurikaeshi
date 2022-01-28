package lib

import (
	_ "embed"
	"encoding/csv"
	"strings"
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
