package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
  var dat []byte
  dat, _ = os.ReadFile("./day5-test-input")
  actual := part1(string(dat))
  expected := "CMZ"
  fmt.Println(expected)
  
  if actual != expected {
    t.Fatal("Error")
  }
}

