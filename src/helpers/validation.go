package helpers

import "strings"

func StringMatch(target string, match string) bool {
	return strings.Contains(target, match)
}

func StringEndWith(target string, match string) bool {
	return strings.HasSuffix(target, match)
}
