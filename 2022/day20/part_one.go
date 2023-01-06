package main

import (
	_ "embed"
	"os"
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}


func part1() int {
	dat, _ := os.ReadFile("./day20-input")
  input := strings.Split(string(dat), "\n")
  input = input[:len(input)-1]

  var tmp [][]int

  for i, num := range input {
    tmp = append(tmp, []int{i, convInt(num)})
  }

  for i := 0; i < len(tmp); i++ {
    for j := 0; j < len(tmp); j++ {
      if tmp[j][0] != i {
        break
      }
    }
    for tmp[0][0] != i {
      tmp_pop := tmp[0]
      tmp = append(tmp[1:], tmp_pop)
    }

    val := tmp[0]
    tmp = tmp[1:]
    to_pop := val[1] % len(tmp)
    for to_pop < 0 {
      to_pop += len(tmp)
    }

    for m := 0; m < to_pop; m++ {
      tmp_pop := tmp[0]
      tmp = append(tmp[1:], tmp_pop)
    }
    tmp = append(tmp, val)

  }
  j := 0
  for i := 0; i < len(tmp); i++ {
    if tmp[i][1] == 0 {
      break
    }
    j++
  }
  return tmp[(j+1000)%len(tmp)][1] + tmp[(j+2000)%len(tmp)][1] + tmp[(j+3000)%len(tmp)][1]
}
