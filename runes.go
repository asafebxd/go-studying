package main

import (
	"unicode/utf8"

	"strings"
)

func countCharacters(text string) int {
	//return len(text) return the number of bytes
	return len([]rune(text))
}

func normalizeName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

func reverseString(text string) string {
	var reversedText strings.Builder

    for len(text) > 0 {
		runeValue, size := utf8.DecodeLastRuneInString(text)
		reversedText.WriteRune(runeValue)
		text = text[:len(text)-size]
	}
    return reversedText.String()
}

// func reverseString(text string) string {
// 	runes := []rune(text)

// 	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
// 		runes[left], runes[right] = runes[right], runes[left]
// 	}

// 	return string(runes)
// }

func runesExercise() {
	
}