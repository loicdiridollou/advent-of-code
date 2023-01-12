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

func MaxInt(ls ...int) int {
	max := ls[0]

	for _, el := range ls {
		if el > max {
			max = el
		}
	}

	return max
}

func findStart(grid [][]string) [2]int {
	for idx_i, i := range grid {
		for idx_j, j := range i {
			if j == "." {
				return [2]int{idx_i, idx_j}
			}
		}
	}
	return [2]int{-1, -1}
}

func splitPath(path string) []string {
	res := make([]string, 0)
	tmp := ""
	for i := 0; i < len(path); i++ {
		if path[i] <= 'Z' && 'A' <= path[i] {
			res = append(res, tmp)
			tmp = ""
			res = append(res, string(path[i]))
		} else {
			tmp = tmp + string(path[i])
      if i == len(path)-1 {
        res = append(res, tmp)
      }
		}
	}

	return res
}

func convertDir(dir string) [2]int {
	var res [2]int
	switch dir {
	case "R":
		res = [2]int{0, 1}
	case "L":
		res = [2]int{0, -1}
	case "U":
		res = [2]int{-1, 0}
	case "D":
		res = [2]int{1, 0}
	}
	return res
}

func moveOnGrid(grid [][]string, start_pos [2]int, steps int, dir string) [2]int {
  rr, cc := convertDir(dir)[0], convertDir(dir)[1]
	for i := 0; i < steps; i++ {
    nr, nc := (start_pos[0]+rr+len(grid))%len(grid), (start_pos[1]+cc+len(grid[0]))%len(grid[0])
    if grid[nr][nc] == "#" {
      break
    } else if grid[nr][nc] == " " {
      tmp := start_pos
      for grid[nr][nc] == " " {
        nr, nc = (start_pos[0]+rr+len(grid))%len(grid), (start_pos[1]+cc+len(grid[0]))%len(grid[0])
        start_pos = [2]int{nr, nc}
      }
      if grid[nr][nc] == "#" {
        nr, nc := (tmp[0]-rr+len(grid))%len(grid), (tmp[1]-cc+len(grid[0]))%len(grid[0])
        start_pos = [2]int{nr, nc}
      } else {
        start_pos = [2]int{nr, nc}
      }

    } else {
      start_pos = [2]int{nr, nc}
    }
  }
	return start_pos
}


func getDest(r int, c int, d int, grid [][]string, D [][]int) (int, int, int) {
  r = (r+D[d][0] + len(grid))%len(grid)
  c = (c+D[d][1]+len(grid[0]))%len(grid[0])
  for grid[r][c] == " " {
    r = (r+D[d][0] + len(grid))%len(grid)
    c = (c+D[d][1]+len(grid[0]))%len(grid[0])
  }
  return r, c, d
}

func part1() int {
	dat, _ := os.ReadFile("./day22-input")
	input := strings.Split(string(dat), "\n\n")
	maze := input[0]
	path := splitPath(strings.Replace(input[1], "\n", "", 1))

	var max int

	for _, el := range strings.Split(string(maze), "\n") {
		if len(el) == 0 {
			continue
		}
		max = MaxInt(max, len(el))
	}

	var grid [][]string

	for _, el := range strings.Split(string(maze), "\n") {
		for len(el) < max {
			el = el + " "
		}
		grid = append(grid, strings.Split(el, ""))
	}

  
  r, c, d := 0, 0, 0
  var rr, cc int
  for grid[r][c] != "." {
    c++
  }
  dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i := 0; i < (len(path)/2)*2; i += 2 {
    for j := 0; j < convInt(path[i]); j++ {
      rr = (r+dirs[d][0]+len(grid))%len(grid)
      cc = (c+dirs[d][1]+len(grid[0]))%len(grid[0])
      if grid[rr][cc] == " " {
        nr, nc, nd := getDest(r, c, d, grid, dirs)
        if grid[nr][nc] == "#" {
          break
        }
        r, c, d = nr, nc, nd
        continue
      } else if grid[rr][cc] == "#" {
        break
      } else {
        r, c  = rr, cc
      }
    }
    if path[i+1] == "L" {
      d = (d+3)%4
    } else {
      d = (d+1)%4
    }
  }
  if len(path) % 2 == 1 {
    for k := 0; k < convInt(path[len(path)-1]); k++ {
      rr = (r+dirs[d][0]+len(grid))%len(grid)
      cc = (c+dirs[d][1]+len(grid[0]))%len(grid[0])
      if grid[rr][cc] == " " {
        nr, nc, nd := getDest(r, c, d, grid, dirs)
        if grid[nr][nc] == "#" {
          break
        }
        r, c, d = nr, nc, nd
        continue
      } else if grid[rr][cc] == "#" {
        break
      } else {
        r, c  = rr, cc
      }
    }
  }

	return 1000 * (r + 1) + 4 * (c + 1) + d 
}
