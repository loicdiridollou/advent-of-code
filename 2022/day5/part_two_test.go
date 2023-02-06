package main

import (
	"os"
	"testing"
)

func Test_part2(t *testing.T) {
  var dat []byte
  dat, _ = os.ReadFile("./day5-test-input")
  actual := part2(string(dat))
  expected := "MCD"
  
  if actual != expected {
    t.Errorf("part2() = %v but expects %v", actual, expected)
  }
}

