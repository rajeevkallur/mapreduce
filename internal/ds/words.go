package ds

import (
	"math"
	"strings"
)

// onesWords holds the English words for the numbers 0 through 19, indexed by value.
var onesWords = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen",
}

// tensWords holds the English words for the multiples of ten from 20 to 90,
// indexed by the tens digit (e.g. tensWords[4] == "forty").
var tensWords = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

// scaleWords holds the English words for each group of three digits, indexed by
// group position (e.g. scaleWords[1] == "thousand", scaleWords[2] == "million").
var scaleWords = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion"}

// NumberToWords converts an integer into its English words representation.
func NumberToWords(n int) string {
	if n == 0 {
		return "zero"
	}

	negative := false
	if n < 0 {
		negative = true
		n = int(math.Abs(float64(n)))
	}

	// Split the number into groups of three digits, least significant first.
	var groups []int
	for n > 0 {
		groups = append(groups, n%1000)
		n /= 1000
	}

	var parts []string
	for i := len(groups) - 1; i >= 0; i-- {
		group := groups[i]
		if group == 0 {
			continue
		}
		words := threeDigitsToWords(group)
		if scaleWords[i] != "" {
			words += " " + scaleWords[i]
		}
		parts = append(parts, words)
	}

	result := strings.Join(parts, " ")
	if negative {
		result = "negative " + result
	}
	return result
}

// threeDigitsToWords converts a number between 1 and 999 into words.
func threeDigitsToWords(n int) string {
	var parts []string

	if n >= 100 {
		parts = append(parts, onesWords[n/100], "hundred")
		n %= 100
	}

	if n >= 20 {
		parts = append(parts, tensWords[n/10])
		n %= 10
		if n > 0 {
			parts = append(parts, onesWords[n])
		}
	} else if n > 0 {
		parts = append(parts, onesWords[n])
	}

	return strings.Join(parts, " ")
}
