package main

func part2(input string) int {
	dic := make(map[string]int)
	for i := 0; i < len(input)-1; i++ {
		if i >= 14 {
			dic[string(input[i-14])] = dic[string(input[i-14])] - 1
		}
		dic[string(input[i])] = dic[string(input[i])] + 1
		if allUnique(dic, 14) {
			return i + 1
		}
	}
	return -1
}
