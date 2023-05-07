package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	var dat []byte
	dat, _ = os.ReadFile("./day16.testinput")
	actual := part1(string(dat))
	var expected int64 = 6

	if actual != expected {
		t.Errorf("part1() = %v but expects %v", actual, expected)
	}
}
