package main

import (
	"fmt"
	"unicode"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

type SkipString struct {
	ind int
	s   string
}

var counter int

func NewSkipString(c int, s string) SkipString {
	return SkipString{c, s}
}

func (t *SkipString) TransformRune(pos int) {
	sl := t.GetValueAsRuneSlice()
	if unicode.IsLetter(sl[pos]) || unicode.IsDigit(sl[pos]) {
		counter++
	} else {
		return
	}

	if counter == 3 && unicode.IsLower(sl[pos]) {
		sl[pos] = unicode.ToUpper(sl[pos])
		counter = 0
	} else if unicode.IsUpper(sl[pos]) {
		sl[pos] = unicode.ToLower(sl[pos])

	}
	t.s = string(sl)
}

func (t *SkipString) GetValueAsRuneSlice() []rune {
	return []rune(t.s)
}

func (t SkipString) String() string {
	return fmt.Sprintf("%v", t.s)
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}
