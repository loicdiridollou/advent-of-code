package main

import "strings"

func part2(input string) string {
	s := strings.Split(input, "\n\n")
	p1 := strings.Split(s[0], "\n")
	p2 := strings.Split(s[1], "\n")
	sl := processStack(p1)
	cmd := processInstructions(p2)
	res := executeCommands(sl, cmd, false)
	return extractTop(res)
}
