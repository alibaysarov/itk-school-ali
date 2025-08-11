package main

import (
	"fmt"
	"testing"
)

func TestGetFuncType(t *testing.T) {
	testCase := struct {
		val    any
		output string
	}{
		val:    "hello",
		output: "string",
	}
	res := GetFuncType(testCase.val)
	fmt.Println(res)
	if res != testCase.output {
		t.Errorf("Test faile with params %s %s", res, testCase.output)
	}
}
