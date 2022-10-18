package helpers

import (
	"encoding/json"
	"regexp"
	"strings"
)

func CamelCaseToSnakeCase(str string) string {
	snake := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(str, "${1}_${2}")
	snake = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func LengthOfString(value string) int {
	return len([]rune(value))
}

func InterfaceToString(value map[string]interface{}) string {
	test, _ := json.Marshal(value)
	return string(test)
}

func MapToJson(value []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(value, &result)
	return result
}

func StringToByte(value string) []byte {
	return []byte(value)
}
