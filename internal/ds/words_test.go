package ds

import "testing"

// TestNumberToWords verifies that NumberToWords converts integers, including
// zero, negative values, and large multi-group numbers, into their expected
// English words representation.
func TestNumberToWords(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"zero", 0, "zero"},
		{"single digit", 7, "seven"},
		{"ten", 10, "ten"},
		{"teen", 13, "thirteen"},
		{"twenty", 20, "twenty"},
		{"tens with ones", 42, "forty two"},
		{"ninety nine", 99, "ninety nine"},
		{"one hundred", 100, "one hundred"},
		{"hundred with remainder", 305, "three hundred five"},
		{"hundred tens ones", 999, "nine hundred ninety nine"},
		{"one thousand", 1000, "one thousand"},
		{"thousand with hundreds", 1234, "one thousand two hundred thirty four"},
		{"thousands no hundreds", 21000, "twenty one thousand"},
		{"million", 1000000, "one million"},
		{"complex million", 1234567, "one million two hundred thirty four thousand five hundred sixty seven"},
		{"skipped group", 1000001, "one million one"},
		{"negative", -5, "negative five"},
		{"negative large", -1234, "negative one thousand two hundred thirty four"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberToWords(tt.n)
			if got != tt.want {
				t.Errorf("NumberToWords(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}

// TestThreeDigitsToWords verifies that threeDigitsToWords converts numbers in
// the range 1 to 999 into their expected English words representation.
func TestThreeDigitsToWords(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "one"},
		{15, "fifteen"},
		{50, "fifty"},
		{55, "fifty five"},
		{100, "one hundred"},
		{110, "one hundred ten"},
		{111, "one hundred eleven"},
		{120, "one hundred twenty"},
		{123, "one hundred twenty three"},
		{900, "nine hundred"},
	}

	for _, tt := range tests {
		got := threeDigitsToWords(tt.n)
		if got != tt.want {
			t.Errorf("threeDigitsToWords(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
