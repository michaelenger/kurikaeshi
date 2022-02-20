package data

import (
	"math/rand"
	"strings"
)

// An item in the registry.
type RegistryItem struct {
	mora  string
	words []*Word
	score int
}

// A word registry which is responsible for picking words to guess.
type Registry struct {
	items []RegistryItem
}

// Register a score change for a word.
func (r Registry) registerScoreChange(word *Word, change int) {
	morae := strings.Split(word.Morae, "")
	for _, mora := range morae {
		for i, item := range r.items {
			if item.mora == mora {
				r.items[i].score += change
				break
			}
		}
	}
}

// Pick a random word from the registry to test.
func (r Registry) PickWord() *Word {
	// TODO: use the first third and sort the list
	// oneThird := int(math.Ceil(float64(len(r.items)) * 0.333))
	index := rand.Intn(len(r.items))
	words := r.items[index].words

	return words[rand.Intn(len(words))]
}

// Register a failed guess of a word.
func (r Registry) RegisterFailure(word *Word) {
	r.registerScoreChange(word, -1)
}

// Register a successful guess of a word.
func (r Registry) RegisterSuccess(word *Word) {
	r.registerScoreChange(word, 1)
}

// Create a registry based on a list of words.
func CreateRegistry(words *[]Word) Registry {
	var moraToWordsMap = make(map[string][]*Word)

	// Pointer fix thanks to: https://tam7t.com/golang-range-and-pointers/
	for i, _ := range *words {
		word := &(*words)[i]
		morae := strings.Split(word.Morae, "")

		for _, mora := range morae {
			words, ok := moraToWordsMap[mora]
			if !ok {
				words = []*Word{}
			}
			moraToWordsMap[mora] = append(words, word)
		}
	}

	var registry Registry
	for mora, words := range moraToWordsMap {
		registry.items = append(registry.items, RegistryItem{mora, words, 0})
	}

	return registry
}
