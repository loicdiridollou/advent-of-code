package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	var dat []byte
	dat, _ = os.ReadFile("./day14.testinput")
	actual := part1(string(dat))
	expected := 1588

	if actual != expected {
		t.Errorf("part1() = %v but expects %v", actual, expected)
	}
}