package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	var dat []byte
	dat, _ = os.ReadFile("./testinput.txt")
	actual := part1(string(dat))
	expected := 1928

	if actual != expected {
		t.Fatal("Error")
	}
}
