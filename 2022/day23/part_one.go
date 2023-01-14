package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	r, c int
}

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func hashPosition(elf Elf) string {
	return fmt.Sprint(elf.r) + "_" + fmt.Sprint(elf.c)
}

func parseInput(input []string) []Elf {
	res := make([]Elf, 0)
	for i, el := range input {
		if len(el) == 0 {
			continue
		}
		for j, chr := range el {
			if string(chr) == "#" {
				res = append(res, Elf{i, j})
			}
		}
	}
	return res
}

func hashElves(elves []Elf) map[string]bool {
	elf_map := make(map[string]bool, 0)

	for _, elf := range elves {
		elf_map[hashPosition(elf)] = true
	}

	return elf_map
}

func nextMove(elf Elf, elf_map map[string]bool, rot int) string {
	r, c := elf.r, elf.c
	nw := hashPosition(Elf{r - 1, c - 1})
	nn := hashPosition(Elf{r - 1, c})
	ne := hashPosition(Elf{r - 1, c + 1})
	ee := hashPosition(Elf{r, c + 1})
	se := hashPosition(Elf{r + 1, c + 1})
	ss := hashPosition(Elf{r + 1, c})
	sw := hashPosition(Elf{r + 1, c - 1})
	ww := hashPosition(Elf{r, c - 1})

	cond := []bool{
		!elf_map[nw] && !elf_map[nn] && !elf_map[ne],
		!elf_map[sw] && !elf_map[ss] && !elf_map[se],
		!elf_map[nw] && !elf_map[ww] && !elf_map[sw],
		!elf_map[ne] && !elf_map[ee] && !elf_map[se],
	}

	dir := []string{nn, ss, ww, ee}

	if !elf_map[nw] && !elf_map[nn] && !elf_map[ne] && !elf_map[ee] && !elf_map[se] && !elf_map[ss] &&
		!elf_map[sw] && !elf_map[ww] {
		return ""
	} else if cond[rot%4] {
		return dir[rot%4]
	} else if cond[(rot+1)%4] {
		return dir[(rot+1)%4]
	} else if cond[(rot+2)%4] {
		return dir[(rot+2)%4]
	} else if cond[(rot+3)%4] {
		return dir[(rot+3)%4]
	}

	return ""
}

func (e Elf) Move(new_pos [2]int) Elf {
	e.r, e.c = new_pos[0], new_pos[1]
	return e
}

func parseMove(s string) [2]int {
	spl := strings.Split(s, "_")
	return [2]int{convInt(string(spl[0])), convInt(string(spl[1]))}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findDim(elves []Elf) int {
	min_r, min_c, max_r, max_c := elves[0].r, elves[0].c, elves[0].r, elves[0].c

	for _, elf := range elves[1:] {
		min_r = Min(min_r, elf.r)
		min_c = Min(min_c, elf.c)
		max_r = Max(max_r, elf.r)
		max_c = Max(max_c, elf.c)
	}

	area := (max_r - min_r + 1) * (max_c - min_c + 1)

	return area - len(elves)
}

func findMoves(elves []Elf, rot int) (map[string]string, map[string]int) {
	moves := make(map[string]string, 0)
	list_move := make(map[string]int, 0)
	elf_map := hashElves(elves)
	for _, elf := range elves {
		if move := nextMove(elf, elf_map, rot); move != "" {
			moves[hashPosition(elf)] = move
			list_move[move] += 1
		}
	}
	return moves, list_move
}

func part1() int {
	dat, _ := os.ReadFile("./day23-input")
	input := strings.Split(string(dat), "\n")

	elves := parseInput(input)
	rot := 0

	for round := 0; round < 10; round++ {
		moves, list_move := findMoves(elves, rot)
		for i, elf := range elves {
			move, moving := moves[hashPosition(elf)]
			if !moving {
				continue
			} else if list_move[move] != 1 {
				continue
			} else {
				elves[i] = elf.Move(parseMove(move))
			}
		}
		rot = (rot + 1) % 4
	}

	return findDim(elves)
}
