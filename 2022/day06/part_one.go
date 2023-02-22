package main

func allUnique(dic map[string]int, num int) bool {
	score := 0
	for _, el := range dic {
		if el != 0 {
			score++
		}
	}
	return (score == num)
}

func part1(input string) int {
	dic := make(map[string]int)
	for i := 0; i < len(input)-1; i++ {
		if i >= 4 {
			dic[string(input[i-4])] = dic[string(input[i-4])] - 1
		}
		dic[string(input[i])] = dic[string(input[i])] + 1
		if allUnique(dic, 4) {
			return i + 1
		}
	}
	return -1
}
