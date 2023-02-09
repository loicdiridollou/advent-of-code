package main

import (
	"os"
	"testing"
)

func Test_part2(t *testing.T) {
  var dat []byte
  dat, _ = os.ReadFile("./day4-test-input")
  actual := part2(string(dat))
  expected := 4
  
  if actual != expected {
    t.Errorf("part2() = %v but expects %v", actual, expected)
  }
}

