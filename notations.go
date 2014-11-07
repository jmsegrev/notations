package notations

import (
	"strings"
	"text/scanner"
)

type Notations map[string][]Notation

type Slice struct {
	Begin int
	End   int
}

type Notation struct {
	Key   rune
	Value string
	Slice Slice
}

// Returns the Notations extracted from the text, uses the distics chars
// as notation indetifier.
func Extract(text string, chars string) Notations {

	notations := make(Notations)

	var s scanner.Scanner
	s.Init(strings.NewReader(text))

	token := s.Scan()
	for token != scanner.EOF {
		if !strings.ContainsRune(chars, token) {
			token = s.Scan()
			continue
		}

		p := s.Pos()
		_ = s.Scan()

		nextIndex := p.Offset + 1
		// skip if next char is: outside text, space, or any of the keys
		if len(text) <= nextIndex || text[p.Offset:nextIndex] == " " ||
			strings.ContainsAny(text[p.Offset:nextIndex], chars) {
			token = s.Scan()
			continue
		}

		index := p.Offset - 1
		// ok if char is at or previous char is space
		if index == 0 || text[index-1:index] == " " {

			notation := Notation{
				Key:   token,
				Value: s.TokenText(),
				Slice: Slice{
					Begin: p.Offset,
					End:   p.Offset + len(s.TokenText()),
				},
			}

			t := string(token)
			if notations[t] == nil {
				notations[t] = []Notation{notation}
			} else {
				notations[t] = append(notations[t], notation)
			}
		}
		token = s.Scan()
	}

	return notations
}
