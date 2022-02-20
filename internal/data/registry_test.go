package data

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestCreateRegistry(t *testing.T) {
	words := []Word{
		{"ひと", "人", "person", "hito"},
		{"ひざ", "膝", "knee", "hiza"},
		{"とり", "鳥", "bird", "tori"},
		// TODO: ensure that it doesn't add the same word twice
	}

	result := CreateRegistry(&words)
	expected := Registry{
		[]RegistryItem{
			{"ひ", []*Word{&words[0], &words[1]}, 0},
			{"と", []*Word{&words[0], &words[2]}, 0},
			{"ざ", []*Word{&words[1]}, 0},
			{"り", []*Word{&words[2]}, 0},
		},
	}

	// TODO: this breaks sometimes, depending on the order of the words
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Unexpected result: %+v != %+v", result, expected)
	}
}

func TestPickWord(t *testing.T) {
	rand.Seed(0) // ensures the same result every time

	word := Word{"ひざ", "膝", "knee", "hiza"}

	registry := Registry{
		[]RegistryItem{
			{"ひ", []*Word{}, 0},
			{"と", []*Word{}, 0},
			{"ざ", []*Word{}, 0},
			{"り", []*Word{}, 0},
			{"お", []*Word{&word}, 0},
			{"く", []*Word{}, 0},
			{"せ", []*Word{}, 0},
			{"し", []*Word{}, 0},
			{"け", []*Word{}, 0},
			{"へ", []*Word{}, 0},
		},
	}
	result := registry.PickWord()
	expected := &word

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Unexpected result: %+v != %+v", result, expected)
	}
}

func TestRegisterFailure(t *testing.T) {
	word := Word{"ひと", "人", "person", "hito"}
	registry := Registry{
		[]RegistryItem{
			{"ひ", []*Word{}, 1},
			{"ざ", []*Word{}, 2},
			{"と", []*Word{}, 2},
			{"り", []*Word{}, 3},
		},
	}
	registry.RegisterFailure(&word)

	expected := Registry{
		[]RegistryItem{
			{"ひ", []*Word{}, 0},
			{"ざ", []*Word{}, 2},
			{"と", []*Word{}, 1},
			{"り", []*Word{}, 3},
		},
	}
	if !reflect.DeepEqual(registry, expected) {
		t.Fatalf("Unexpected registry: %+v != %+v", registry, expected)
	}
}

func TestRegisterSuccess(t *testing.T) {
	word := Word{"ひと", "人", "person", "hito"}
	registry := Registry{
		[]RegistryItem{
			{"ひ", []*Word{}, 1},
			{"ざ", []*Word{}, 2},
			{"と", []*Word{}, 2},
			{"り", []*Word{}, 3},
		},
	}
	registry.RegisterSuccess(&word)

	expected := Registry{
		[]RegistryItem{
			{"ひ", []*Word{}, 2},
			{"ざ", []*Word{}, 2},
			{"と", []*Word{}, 3},
			{"り", []*Word{}, 3},
		},
	}
	if !reflect.DeepEqual(registry, expected) {
		t.Fatalf("Unexpected registry: %+v != %+v", registry, expected)
	}
}
