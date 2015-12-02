package a1strings

import (
	"regexp"
	"strings"
)

// Rpad -
//   Padding space to right.
func Rpad(s string, maxlen int) string {
	if maxlen > len(s) {
		return s + strings.Repeat(" ", maxlen-len(s))
	}
	return s
}

// Lpad -
//   Padding space to left.
func Lpad(s string, maxlen int) string {
	if maxlen > len(s) {
		return strings.Repeat(" ", maxlen-len(s)) + s
	}
	return s
}

// RpadFontWidth -
//   Padding space to right.
func RpadFontWidth(s string, maxlen int) string {
	if maxlen > LenFontWidth(s) {
		return s + strings.Repeat(" ", maxlen-LenFontWidth(s))
	}
	return s
}

// LpadFontWidth -
//   Padding space to left.
func LpadFontWidth(s string, maxlen int) string {
	if maxlen > LenFontWidth(s) {
		return strings.Repeat(" ", maxlen-LenFontWidth(s)) + s
	}
	return s
}

// LenFontWidth -
//   Return font width.
//   This function treat Half wide katakana lazy. It counted 2 :P
func LenFontWidth(s string) int {
	lenc := len([]rune(s))
	lenb := len(s)

	if lenc == lenb {
		return lenc
	} else {
		return lenc + ((lenb - lenc) / 2)
	}
}

// Snake2Camel -
//   Change snake case string to lower camel case.
func Snake2Camel(s string) string {
	re, _ := regexp.Compile("_([a-z0-9])")
	cb := func(s2 string) string {
		m := re.FindStringSubmatch(s2)
		return strings.ToUpper(m[1])
	}
	return re.ReplaceAllStringFunc(strings.ToLower(s), cb)
}
