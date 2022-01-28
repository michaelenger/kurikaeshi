package challenge

import "testing"

func TestSanitize(t *testing.T) {
	var tests = map[string]string{
		"hitsuji":         "hitsuji",
		"te-no-hira":      "tenohira",
		"kanzō":           "kanzou",
		"sūji":            "suuji",
		"sobo (obaa-san)": "sobo",
	}

	for test, expected := range tests {
		result := sanitize(test)
		if result != expected {
			t.Fatalf("Unexpected result: %v != %v", result, expected)
		}
	}
}
