package wordler

import (
	"errors"
	"strings"
)

var (
	ErrIndexOverflow = errors.New("index >= len(s)")
)

type Filter interface {
	Check(string) bool
}

type FilterList []Filter

func (l FilterList) Run(words []string) []string {
	result := []string{}
	for _, word := range words {
		valid := true
		for i := range l {
			if !l[i].Check(word) {
				valid = false
				break
			}
		}
		if valid {
			result = append(result, word)
		}
	}
	return result
}

type FilterContain struct {
	Target string
}

func (f FilterContain) Check(s string) bool {
	return strings.Contains(s, f.Target)
}

type FilterNotContain struct {
	Target string
}

func (f FilterNotContain) Check(s string) bool {
	return !strings.Contains(s, f.Target)
}

type FilterAt struct {
	Target byte
	Index  int
}

func (f FilterAt) Check(s string) bool {
	return f.Index < len(s) && f.Index >= 0 && s[f.Index] == f.Target
}

type FilterNotAt struct {
	Target byte
	Index  int
}

func (f FilterNotAt) Check(s string) bool {
	return f.Index >= len(s) || f.Index < 0 || s[f.Index] != f.Target
}
