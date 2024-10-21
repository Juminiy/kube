package util

import "testing"

func TestCamel2Snake(t *testing.T) {
	t.Log(Camel2Snake("camel"))
	t.Log(Camel2Snake("camelToSnake"))
	t.Log(Camel2Snake("camel2Snake"))
	t.Log(Camel2Snake(""))
	t.Log(Camel2Snake("camel_to_snake"))
	t.Log(Camel2Snake("______________"))
	t.Log(Camel2Snake("CAMEL_TO_SNAKE"))
	t.Log(Camel2Snake("CAMELTOSNAKE"))

	t.Log(Snake2Camel("camel"))
	t.Log(Snake2Camel("camelToSnake"))
	t.Log(Snake2Camel("camel2Snake"))
	t.Log(Snake2Camel(""))
	t.Log(Snake2Camel("camel_to_snake"))
	t.Log(Snake2Camel("______________"))
	t.Log(Snake2Camel("CAMEL_TO_SNAKE"))
	t.Log(Snake2Camel("CAMELTOSNAKE"))
}
