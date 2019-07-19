package string_utils

import (
	"strings"
	"fmt"
)

func isPalindrome(text string) bool {
	if len(text) <= 1 {
		return true
	}

	if len(text) == 2 && text[0] == text[1] {
		return true
	}

	if text[0] != text[len(text)-1] {
		return false
	}

	return isPalindrome(text[1 : len(text)-1])
}

func IsPalindrome(text string) bool {
	fmt.Println("keep calling strings utils library")
	return isPalindrome(strings.Replace(text, " ", "", -1))
}
