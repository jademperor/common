package utils

import (
	"strings"
)

// StringArray type of []string to get array form cli proc
type StringArray []string

// String get all value in StringArray
func (s *StringArray) String() string {
	return strings.Join(*s, "/")
}

// Set value into StringArray
func (s *StringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}
