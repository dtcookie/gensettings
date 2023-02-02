package camel

import (
	"strings"
	"unicode"
)

const (
	nonletter  = 0
	letter     = 1
	underscore = 2
	digit      = 3
	lower      = 4
	upper      = 5
)

func Camel(s string) string {
	if len(s) == 0 {
		return s
	}
	res := ""
	lastKind := 0
	for _, ch := range s {
		kind := 0
		if unicode.IsLetter(ch) {
			kind = letter
			if len(res) == 0 || lastKind != letter {
				res = res + string(unicode.ToUpper(ch))
			} else {
				res = res + string(unicode.ToLower(ch))
			}
		} else if unicode.IsDigit(ch) {
			kind = digit
			res = res + string(ch)
		} else if ch == '_' {
			kind = underscore
		} else {
			kind = nonletter
		}
		lastKind = kind
	}
	return res
}

func Strip(s string) string {
	if len(s) == 0 {
		return s
	}
	res := ""
	lastKind := 0
	for _, ch := range s {
		kind := 0
		if unicode.IsLetter(ch) {
			if unicode.IsUpper(ch) {
				kind = upper
			} else {
				kind = lower
			}
		} else {
			kind = nonletter
		}

		if (len(res) > 0) && (((lastKind == nonletter) && (kind != nonletter)) || ((lastKind == upper) && (kind == nonletter)) || ((lastKind == lower) && (kind != lower))) {
			res = res + "_"
		}
		res = res + string(unicode.ToLower(ch))
		lastKind = kind
	}
	for strings.Contains(res, "__") {
		res = strings.ReplaceAll(res, "__", "_")
	}
	return res
}
