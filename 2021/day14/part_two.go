package main

func part2(input string) int {
	formula, rules := prepareData(input)

	dic := buildMap(formula, rules)
	counter := Counter(formula)
	rounds := 40

	for num_round := 0; num_round < rounds; num_round++ {
		new_dic := copyMap(dic)
		for _, rule := range rules {
			if val, ok := dic[rule.pattern]; ok && val > 0 {
				left := string(rule.pattern[0]) + rule.addition
				right := rule.addition + string(rule.pattern[1])

				new_dic[rule.pattern] -= val
				new_dic[left] += val
				new_dic[right] += val
				if _, ok := counter[rule.addition]; !ok {
					counter[rule.addition] = 0
				}
				counter[rule.addition] += val
			}
		}
		dic = new_dic
	}

	min, max := getMinMax(counter)

	return max - min
}
