package main

import (
	_ "embed"
)

type Coord struct {
	x, y int
}

func (c Coord) Equals(cc Coord) bool {
	return c.x == cc.x && c.y == cc.y
}

type Floor struct {
	lim []Coord
}

type Shape struct {
	coord []Coord
}

func getPiece(num int, top int) Shape {
	var shape Shape
	switch num {
	case 0:
		shape = Shape{[]Coord{{2, top}, {3, top}, {4, top}, {5, top}}}
	case 1:
		shape = Shape{[]Coord{{2, top + 1}, {3, top + 2}, {3, top + 1}, {3, top}, {4, top + 1}}}
	case 2:
		shape = Shape{[]Coord{{2, top}, {3, top}, {4, top}, {4, top + 1}, {4, top + 2}}}
	case 3:
		shape = Shape{[]Coord{{2, top}, {2, top + 1}, {2, top + 2}, {2, top + 3}}}
	case 4:
		shape = Shape{[]Coord{{2, top}, {3, top}, {2, top + 1}, {3, top + 1}}}
	default:
		shape = Shape{}
	}
	return shape
}

func (f Floor) Intersect(shape Shape) bool {
	for _, c1 := range f.lim {
		for _, c2 := range shape.coord {
			if c1.Equals(c2) {
				return true
			}
		}
	}
	return false
}

func (f Floor) Update(piece Shape) Floor {
	f.lim = append(f.lim, piece.coord...)
	return f
}

func moveLeft(piece Shape) Shape {
	for _, coord := range piece.coord {
		if coord.x == 0 {
			return piece
		}
	}
	new_piece := Shape{}
	for _, coord := range piece.coord {
		new_piece.coord = append(new_piece.coord, Coord{coord.x - 1, coord.y})
	}
	return new_piece
}

func moveRight(piece Shape) Shape {
	for _, coord := range piece.coord {
		if coord.x == 6 {
			return piece
		}
	}
	new_piece := Shape{}
	for _, coord := range piece.coord {
		new_piece.coord = append(new_piece.coord, Coord{coord.x + 1, coord.y})
	}
	return new_piece
}

func moveUp(piece Shape) Shape {
	new_piece := Shape{}
	for _, coord := range piece.coord {
		new_piece.coord = append(new_piece.coord, Coord{coord.x, coord.y + 1})
	}
	return new_piece
}

func moveDown(piece Shape) Shape {
	new_piece := Shape{}
	for _, coord := range piece.coord {
		new_piece.coord = append(new_piece.coord, Coord{coord.x, coord.y - 1})
	}
	return new_piece
}

func Max(lim []Coord) int {
	max := 0
	for _, coord := range lim {
		if max < coord.y {
			max = coord.y
		}
	}
	return max
}

func part1(input string) int {
	s := input[:len(input)-1]

	tab := make([]Coord, 7)
	for i := 0; i < 7; i++ {
		tab[i] = Coord{i, 0}
	}
	floor := Floor{tab}
	L := 2022
	i := 0
	t := 0
	top := 0
	added := 0

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
				break
			}
		}
		t += 1
	}
	return top + added
}
