package ui

import "testing"

func TestNormalizeKey(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"q", "q"},
		{"й", "q"},
		{"Й", "q"},
		{"p", "p"},
		{"з", "p"},
		{"З", "p"},
		{"s", "s"},
		{"ы", "s"},
		{"Ы", "s"},
		{"o", "o"},
		{"щ", "o"},
		{"Щ", "o"},
		{"a", "a"},
		{"ф", "a"},
		{"Ф", "a"},
		{"n", "n"},
		{"т", "n"},
		{"Т", "n"},
		{"i", "i"},
		{"ш", "i"},
		{"Ш", "i"},
		{"k", "k"},
		{"л", "k"},
		{"j", "j"},
		{"о", "j"},
	}

	for _, tt := range tests {
		actual := NormalizeKey(tt.input)
		if actual != tt.expected {
			t.Errorf("NormalizeKey(%q) = %q; want %q", tt.input, actual, tt.expected)
		}
	}
}
