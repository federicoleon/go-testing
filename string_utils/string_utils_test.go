package string_utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestMain(m *testing.M) {
	//Do whatever you need to do before running the test cases.

	os.Exit(m.Run())
}

func TestIsPalindromeLenOne(t *testing.T) {
	assert.True(t, IsPalindrome("a"), "single character should be a palindrome")
}

func TestIsPalindromeTwoSameCharacters(t *testing.T) {
	assert.True(t, IsPalindrome("aa"), "two same characters should be a palindrome")
}

func TestIsPalindromeNotPalindrome(t *testing.T) {
	assert.False(t, IsPalindrome("abc"), "not palindrome should return false")
}

func TestIsPalindromeNoError(t *testing.T) {
	assert.True(t, IsPalindrome("anitalavalatina"), "palindrome should return true")
}

func TestIsPalindromeWithSpacesNoReplaces(t *testing.T) {
	assert.True(t, IsPalindrome("anita lava la tina"), "palindrome should return true")
}

func TestIsPalindromeWithSpacesNoError(t *testing.T) {
	if IsPalindrome("anita lava la tina") != true {
		t.Error("palindrome should return true")
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		IsPalindrome("anita lava la tina")
	}
}
