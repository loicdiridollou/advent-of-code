package main

import (
	_ "embed"
	"os"
	"strings"
)

func BFS(start [2]int, elev [][]string) int {
  queue := [][3]int{{start[0], start[1], 0}}
  seen := make(map[[2]int]bool, 0)
  dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
  for len(queue) > 0 {
    front := queue[0]
    queue = queue[1:]

    if !seen[[2]int{front[0], front[1]}] {
      seen[[2]int{front[0], front[1]}] = true
      r, c := front[0], front[1]
      if elev[r][c] == "E" {
        return front[2]
      } else {
        chr := elev[r][c]
        if chr == "S" {
          chr = "a"
        } else if chr == "E" {
          chr = "z"
        }
        height := int(chr[0])
        for _, dir := range dirs {
          nr, nc := r + dir[0], c + dir[1]
          if 0 <= nr && nr < len(elev) && 0 <= nc && nc < len(elev[0]) {
            chr = elev[nr][nc]
            if chr == "S" {
              chr = "a"
            } else if chr == "E" {
              chr = "z"
            }
            height2 := int(chr[0])
            if height2 <= height + 1 {
              queue = append(queue, [3]int{nr, nc, front[2]+1})
            }
          }
        }
      }
    }
  }
  return -1
}
  
func part1() int {
	dat, _ := os.ReadFile("./day12-input")
	s := strings.Split(string(dat), "\n")
	res := make([][]string, 0)
	start_at := "S"
  start_pos := make([][2]int, 0)

	for idx, u := range s[:len(s)-1] {
		spl := strings.Split(u, "")
		if strings.Contains(u, start_at) {
			start_pos = append(start_pos, [2]int{idx, strings.Index(u, start_at)})
		}
		res = append(res, spl)
	}

	return BFS(start_pos[0], res)
}
