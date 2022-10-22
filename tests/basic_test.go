package tests

import (
	"microservice/src/helpers"
	"microservice/src/services"
	"microservice/src/utils/environment"
	"testing"
)

func TestBasicAdd(t *testing.T) {
	a, b := 6, 3 // change this value

	hope := a + b
	environment.Get()
	response := services.BasicAdd(uint64(a), uint64(b))
	test1 := helpers.InterfaceToString(response.Render)
	if helpers.StringMatch(test1, "error") {
		t.Errorf("Hmmmm : %v\n", test1)
	} else {
		test2 := helpers.MapToJson(helpers.StringToByte(test1))
		if int64(test2["result"].(float64)) == int64(hope) {
			t.Log(response.Render)
		} else {
			t.Errorf("this is not %v\n", hope)
		}
	}
}

func TestBasicMultiply(t *testing.T) {
	a, b := 6, 3 // change this value

	hope := a * b
	environment.Get()
	response := services.BasicMultiply(uint64(a), uint64(b))
	test1 := helpers.InterfaceToString(response.Render)
	if helpers.StringMatch(test1, "error") {
		t.Errorf("Hmmmm : %v\n", test1)
	} else {
		test2 := helpers.MapToJson(helpers.StringToByte(test1))
		if int64(test2["result"].(float64)) == int64(hope) {
			t.Log(response.Render)
		} else {
			t.Errorf("this is not %v\n", hope)
		}
	}
}
