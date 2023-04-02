package main

import (
	_ "embed"
	"fmt"
	"sort"
)

func signature(r Floor) string {
	maxY := Max(r.lim)
	s := make([][2]int, 0)
	for _, coord := range r.lim {
		if maxY-coord.y <= 30 {
			s = append(s, [2]int{coord.x, maxY - coord.y})
		}
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i] == s[j] {
			return s[i][1] < s[j][1]
		}
		return s[i][0] < s[j][0]
	})
	str := ""
	for _, el := range s {
		str += fmt.Sprint(el[0]) + "_" + fmt.Sprint(el[1]) + "_"
	}
	return str[:len(str)-1]
}

func part2(input string) int {
	s := input[:len(input)-1]

	tab := make([]Coord, 7)
	for i := 0; i < 7; i++ {
		tab[i] = Coord{i, 0}
	}
	floor := Floor{tab}
	L := 1000000000000
	i := 0
	t := 0
	top := 0
	SEEN := make(map[string][2]int, 0)
	added := 0

	var found bool
	for t < L {
		piece := getPiece(t%5, top+4)
		for {
			if string(s[i]) == "<" {
				piece = moveLeft(piece)
				if floor.Intersect(piece) {
					piece = moveRight(piece)
				}
			} else {
				piece = moveRight(piece)
				if floor.Intersect(piece) {
					piece = moveLeft(piece)
				}
			}
			i = (i + 1) % len(s)
			piece = moveDown(piece)

			if floor.Intersect(piece) {
				piece = moveUp(piece)
				floor = floor.Update(piece)
				top = Max(floor.lim)
				tmpSR := fmt.Sprint(i) + "__" + fmt.Sprint(t%5) + "__" + signature(floor)

				_, found = SEEN[tmpSR]
				if found {
					oldt, oldy := SEEN[tmpSR][0], SEEN[tmpSR][1]
					dy := top - oldy
					dt := t - oldt
					amt := (L - t) / dt
					added += amt * dy
					t += amt * dt
				}
				SEEN[tmpSR] = [2]int{t, top}
				break
			}
		}
		t += 1
	}
	return top + added
}
