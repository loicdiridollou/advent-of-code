package main

import (
	"fmt"
	"os"
	"strings"
)

func gameResult(p1 string, p2 string) string {
	res := ""
	if p1 == "A" {
		if p2 == "X" {
			return "Tie"
		} else if p2 == "Y" {
			res = "Win"
		} else if p2 == "Z" {
			res = "Lose"
		}
	} else if p1 == "B" {
		if p2 == "X" {
			res = "Lose"
		} else if p2 == "Y" {
			res = "Tie"
		} else if p2 == "Z" {
			res = "Win"
		}
	} else if p1 == "C" {
		if p2 == "X" {
			res = "Win"
		} else if p2 == "Y" {
			res = "Lose"
		} else if p2 == "Z" {
			res = "Tie"
		}
	}
	return res
}

func adaptPlay(opp string, res string) string {
	play := ""
	if opp == "A" {
		if res == "X" {
			play = "Z"
		} else if res == "Y" {
			play = "X"
		} else if res == "Z" {
			play = "Y"
		}
	} else if opp == "B" {
		if res == "X" {
			play = "X"
		} else if res == "Y" {
			play = "Y"
		} else if res == "Z" {
			play = "Z"
		}
	} else {
		if res == "X" {
			play = "Y"
		} else if res == "Y" {
			play = "Z"
		} else if res == "Z" {
			play = "X"
		}
	}
	return play
}

func part1() int {
	dat, _ := os.ReadFile("./day2-input")
	s := strings.Split(string(dat), "\n")
	res := ""
	score := 0
	for i := 0; i < len(s)-1; i++ {
		res = gameResult(string(s[i][0]), string(s[i][2]))
		if res == "Win" {
			score += 6
		} else if res == "Tie" {
			score += 3
		}
		if string(s[i][2]) == "X" {
			score += 1
		} else if string(s[i][2]) == "Y" {
			score += 2
		} else {
			score += 3
		}
	}
	return score
}

func part2() int {
	dat, _ := os.ReadFile("./day2-input")
	s := strings.Split(string(dat), "\n")
	res := ""
	play := ""
	score := 0
	for i := 0; i < len(s)-1; i++ {
		play = adaptPlay(string(s[i][0]), string(s[i][2]))
		res = gameResult(string(s[i][0]), play)
		if res == "Win" {
			score += 6
		} else if res == "Tie" {
			score += 3
		}
		if play == "X" {
			score += 1
		} else if play == "Y" {
			score += 2
		} else {
			score += 3
		}
	}
	return score
}

func main() {
	// fmt.Println(part1())
	fmt.Println(part2())
	fmt.Println("DONE")
}
