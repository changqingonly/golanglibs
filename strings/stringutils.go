package golang_strings

import "strings"

// a为空, 返回b, 否则返回a
func StrGetOr(a, b string) string {
	if a == "" {
		return b
	}
	return a
}

// a去空格后是否为空, 空则返回b
func StrGetTrimOr(a, b string) string {
	if a = strings.TrimSpace(a); a == "" {
		return b
	}
	return a
}
