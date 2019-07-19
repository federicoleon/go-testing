package services

import (
	"github.com/federicoleon/go-testing/string_utils"
)

var (
	StringsService stringServiceInterface
)

func init() {
	StringsService = &stringsService{}
}

type stringsService struct {
	Name string
}

type stringServiceInterface interface {
	IsPalindrome(someText string) bool
}

func (s *stringsService) IsPalindrome(text string) bool {
	return string_utils.IsPalindrome(text)
}

func (s *stringsService) Method1(text string) bool {
	return string_utils.IsPalindrome(text)
}

func (s *stringsService) Method2(text string) bool {
	return string_utils.IsPalindrome(text)
}
