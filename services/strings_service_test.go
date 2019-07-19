package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	isPalindromeFunction func(t string) bool
)

type serviceMock struct{}

func (s *serviceMock) IsPalindrome(t string) bool {
	return isPalindromeFunction(t)
}

func TestIsPalindromeWithMock(t *testing.T) {
	StringsService = &serviceMock{}

	isPalindromeFunction = func(t string) bool {
		return false
	}
	assert.False(t, StringsService.IsPalindrome("ana"))

	isPalindromeFunction = func(t string) bool {
		return false
	}
	assert.False(t, StringsService.IsPalindrome("ana"))
}

func TestIsPalindromeWithoutMock(t *testing.T) {
	StringsService = &stringsService{}
	assert.True(t, StringsService.IsPalindrome("ana"))
	assert.False(t, StringsService.IsPalindrome("fede"))
}
