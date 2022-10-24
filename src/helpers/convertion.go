package helpers

import (
	"encoding/json"
	"fmt"
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
	out, _ := json.Marshal(value)
	return string(out)
}
func StringToJson(jsonString string) map[string]interface{} {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		fmt.Printf("ERROR: convert string to json, %v\n", err)
	}
	return jsonMap
}

func MapToJson(value []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(value, &result)
	return result
}

func StringToByte(value string) []byte {
	return []byte(value)
}
