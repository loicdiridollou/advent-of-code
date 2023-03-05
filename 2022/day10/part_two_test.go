package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_part2(t *testing.T) {
	var dat []byte
	dat, _ = os.ReadFile("./day10.testinput")
	actual := part2(string(dat))
	expected := [6][]string{}

	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		t.Errorf("part2() = %v but expects %v", actual, expected)
	}
}
