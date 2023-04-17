package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	var dat []byte
	dat, _ = os.ReadFile("./day03.testinput")
	actual := part1(string(dat))
	expected := 198
	fmt.Println(actual)

	if actual != expected {
		t.Errorf("part1() = %v but expects %v", actual, expected)
	}
}
