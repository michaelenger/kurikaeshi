package colors

import (
	"testing"
)

func TestBlue(t *testing.T) {
	result := Blue("text")
	expected := "\033[34mtext\033[0m"
	if result != expected {
		t.Fatalf("Unexpected result: %v != %v", result, expected)
	}
}

func TestCyan(t *testing.T) {
	result := Cyan("text")
	expected := "\033[36mtext\033[0m"
	if result != expected {
		t.Fatalf("Unexpected result: %v != %v", result, expected)
	}
}

func TestGreen(t *testing.T) {
	result := Green("text")
	expected := "\033[32mtext\033[0m"
	if result != expected {
		t.Fatalf("Unexpected result: %v != %v", result, expected)
	}
}

func TestPurple(t *testing.T) {
	result := Purple("text")
	expected := "\033[35mtext\033[0m"
	if result != expected {
		t.Fatalf("Unexpected result: %v != %v", result, expected)
	}
}

func TestRed(t *testing.T) {
	result := Red("text")
	expected := "\033[31mtext\033[0m"
	if result != expected {
		t.Fatalf("Unexpected result: %v != %v", result, expected)
	}
}

func TestYellow(t *testing.T) {
	result := Yellow("text")
	expected := "\033[33mtext\033[0m"
	if result != expected {
		t.Fatalf("Unexpected result: %v != %v", result, expected)
	}
}

func TestWhite(t *testing.T) {
	result := White("text")
	expected := "\033[37mtext\033[0m"
	if result != expected {
		t.Fatalf("Unexpected result: %v != %v", result, expected)
	}
}
