package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
  var dat []byte
  dat, _ = os.ReadFile("./day4-test-input")
  actual := part1(string(dat))
  expected := 2
  
  if actual != expected {
    t.Fatal("Error")
  }
}

