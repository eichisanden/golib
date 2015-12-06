// a1string is utilities for string
package a1strings

import (
	"regexp"
	"strings"
)

// Rpad pads spaces to right end until it reaches the specific byte length.
func Rpad(s string, length int) string {
	if length > len(s) {
		return s + strings.Repeat(" ", length - len(s))
	}
	return s
}

// Lpad pads spaces to right end until it reaches the specific byte length.
func Lpad(s string, length int) string {
	if length > len(s) {
		return strings.Repeat(" ", length-len(s)) + s
	}
	return s
}

// RpadFontWidth pads spaces to right end until it reaches the specific font-wide length.
func RpadFontWidth(s string, length int) string {
	if length > LenFontWidth(s) {
		return s + strings.Repeat(" ", length-LenFontWidth(s))
	}
	return s
}

// LpadFontWidth pads spaces to right end until it reaches the specific font-wide length.
func LpadFontWidth(s string, length int) string {
	if length > LenFontWidth(s) {
		return strings.Repeat(" ", length-LenFontWidth(s)) + s
	}
	return s
}

// LenFontWidth return display font width.
// This function treats half wide katakana lazy(Correct return value is 1 but this function returns 2)
func LenFontWidth(s string) int {
	lenc := len([]rune(s))
	lenb := len(s)

	if lenc == lenb {
		return lenc
	} else {
		return lenc + ((lenb - lenc) / 2)
	}
}

// Snake2Camel returns lower camel case string from snake case string.
func Snake2Camel(s string) string {
	re, _ := regexp.Compile("_([a-z0-9])")
	cb := func(s2 string) string {
		m := re.FindStringSubmatch(s2)
		return strings.ToUpper(m[1])
	}
	return re.ReplaceAllStringFunc(strings.ToLower(s), cb)
}
