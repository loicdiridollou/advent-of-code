package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Block struct {
	xa  float64
	xb  float64
	ya  float64
	yb  float64
	xpr float64
	ypr float64
}

func (b Block) solve() (float64, float64) {
	det := (b.xa * b.yb) - (b.xb * b.ya)
	if det == 0 {
		return -1, -1
	}
	return ((b.xpr * b.yb) - (b.ypr * b.ya)) / det, ((b.xa * b.ypr) - (b.xb * b.xpr)) / det
}

var re = regexp.MustCompile(
	`Button A: X\+(\d*),.Y\+(\d*)\nButton B: X\+(\d*),.Y\+(\d*)\nPrize: X\=(\d+), Y\=(\d+)`,
)

func to_float(str string) float64 {
	val, _ := strconv.ParseFloat(str, 64)
	return val
}

func parse_block(block string, ext int) Block {
	res := re.FindStringSubmatch(block)
	return Block{
		to_float(res[1]),
		to_float(res[2]),
		to_float(res[3]),
		to_float(res[4]),
		to_float(res[5]) + float64(ext),
		to_float(res[6]) + float64(ext),
	}
}

// part1 function
func part1(input string) float64 {
	s := strings.Split(input, "\n\n")

	count := 0.
	for _, block := range s {
		a, b := parse_block(block, 0).solve()
		if 0 < a && a <= 100 && math.Floor(a) == a && 0 < b && b <= 100 && math.Floor(b) == b {
			count += 3*a + b
		}
	}

	return count
}
