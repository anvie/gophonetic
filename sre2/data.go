package sre2

import (
	"container/vector"
	"unicode"
)

// RuneFilter is a unique method signature for matching true/false over a given
// unicode rune.
type RuneFilter func(rune int) bool

// Generate a RuneFilter matching a single rune.
func matchRune(to_match int) RuneFilter {
	return func(rune int) bool {
		return rune == to_match
	}
}

// Generate a RuneFilter matching a range of runes, assumes from <= to.
func matchRuneRange(from int, to int) RuneFilter {
	return func(rune int) bool {
		return rune >= from && rune <= to
	}
}

// Generate a RuneFilter matching a valid Unicode class. If no matching classes
// are found, then this method will return nil.
// Note that if just a single character is given, Categories will be searched
// for this as a prefix (so that 'N' will match 'Nd', 'Nl', 'No' etc).
func matchUnicodeClass(class string) RuneFilter {
	found := false
	var match vector.Vector
	if len(class) == 1 {
		// A single character is a shorthand request for any category starting with this.
		for key, r := range unicode.Categories {
			if key[0] == class[0] {
				found = true
				match.Push(r)
			}
		}
	} else {
		// Search for the unicode class name inside cats/props/scripts.
		options := []map[string]*unicode.RangeTable{
			unicode.Categories, unicode.Properties, unicode.Scripts}
		for _, option := range options {
			if r, ok := option[class]; ok {
				found = true
				match.Push(r)
			}
		}
	}

	if found {
		return func(rune int) bool {
			for _, raw := range match {
				r, _ := raw.(*unicode.RangeTable)
				if unicode.Is(r, rune) {
					return true
				}
			}
			return false
		}
	}
	return nil
}

// Generate and return a new, inverse RuneFilter from the argument.
func (r RuneFilter) not() RuneFilter {
	return func(rune int) bool {
		return !r(rune)
	}
}

// Generate and return a new RuneFilter, which ignores case, from the argument.
func (r RuneFilter) ignoreCase() RuneFilter {
	return func(rune int) bool {
		return r(unicode.ToLower(rune)) || r(unicode.ToUpper(rune))
	}
}
