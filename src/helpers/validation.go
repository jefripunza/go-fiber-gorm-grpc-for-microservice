package helpers

import "strings"

func StringMatch(target string, match string) bool {
	return strings.Contains(target, match)
}
