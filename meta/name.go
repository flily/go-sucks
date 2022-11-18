package meta

import (
	"unicode"
)

func getFirstRune(name string) rune {
	if len(name) <= 0 {
		return 0
	}

	var first rune
	for _, c := range name {
		first = c
		break
	}

	return first
}

// IsExportedName checks if the field name of a struct is exported.
func IsExportedName(name string) bool {
	first := getFirstRune(name)

	// Rob Pike said, for Go 2 (can't do it before then), change the definition
	// to "lower case letters and _ are package local; all else is exported",
	// to support exported names in uncases languages like Japanese.
	//
	// see: https://github.com/golang/go/issues/5763#issuecomment-66081539

	// return !unicode.IsLower(first) && first != '_'
	return unicode.IsUpper(first)
}
