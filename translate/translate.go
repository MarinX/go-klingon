package translate

import (
	"fmt"
	"strings"
)

// Translate model
type Translate string

// New creates transatable method
func New(txt string) Translate {
	return Translate(txt)
}

// Klingon returns string in hex
// Error is thrown on unavailable characters
func (t Translate) Klingon() (string, error) {
	if len(t) <= 0 {
		return "", fmt.Errorf("Cannot translate empty string")
	}
	str := []string{}

	for i := 0; i < len(t); i++ {
		found := false

		// handle special cases for tlh, ch, gh and ng
		// first try 3 chars, then 2 chars and lastly 1 char
		for j := 2; j >= 0; j-- {
			if val, ok := t.search(i, i+j); ok {
				i = i + j
				str = append(str, val)
				found = true
				break
			}
		}

		// character not found, stop!
		if !found {
			return "", fmt.Errorf("%c is not a valid klingon character", t[i])
		}
	}

	return strings.Join(str, " "), nil
}

// search for character from hex map using index and position
func (t Translate) search(index, position int) (string, bool) {
	if position < len(t) {
		substr := string(t[index:(1 + position)])
		if hexMap[substr] != "" {
			return hexMap[substr], true
		}

		// handle for lower cases
		substr = strings.ToLower(substr)
		if hexMap[substr] != "" {
			return hexMap[substr], true
		}

		// handle for upper case
		substr = strings.ToUpper(substr)
		if hexMap[substr] != "" {
			return hexMap[substr], true
		}
	}

	return "", false
}
