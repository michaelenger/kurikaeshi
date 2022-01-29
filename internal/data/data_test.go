package data

import (
	"testing"
)

func TestLoadHiragana(t *testing.T) {
	result, err := LoadHiragana()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedLength := 428
	if len(result) != expectedLength {
		t.Fatalf("Incorrect result length %v != %v", len(result), expectedLength)
	}

	expectedEntry := Word{
		"くりかえし",
		"繰り返し",
		"repetition",
		"kurikaeshi",
	}
	if result[0] != expectedEntry {
		t.Fatalf("Unexpected first result: %+v", result[0])
	}
}

func TestLoadKatakana(t *testing.T) {
	result, err := LoadKatakana()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedLength := 73
	if len(result) != expectedLength {
		t.Fatalf("Incorrect result length %v != %v", len(result), expectedLength)
	}

	expectedEntry := Word{
		"エンジニア",
		"-",
		"engineer",
		"enjinia",
	}
	if result[0] != expectedEntry {
		t.Fatalf("Unexpected first result: %+v", result[0])
	}
}

func TestSanitize(t *testing.T) {
	var tests = map[string]string{
		"hitsuji":         "hitsuji",
		"te-no-hira":      "tenohira",
		"kanzō":           "kanzou",
		"sūji":            "suuji",
		"sobo (obaa-san)": "sobo",
	}

	for test, expected := range tests {
		result := Sanitize(test)
		if result != expected {
			t.Fatalf("Unexpected result: %v != %v", result, expected)
		}
	}
}
