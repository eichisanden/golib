package a1string

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

// Snake2Camel -
//   Change snake case string to lower camel case.
//   Parameter string must be lower case string (ex. user_id).
func Snake2Camel(s string) string {
	re, _ := regexp.Compile("_([a-z0-9])")
	cb := func(s2 string) string {
		m := re.FindStringSubmatch(s2)
		return strings.ToUpper(m[1])
	}
	return re.ReplaceAllStringFunc(s, cb)
}
