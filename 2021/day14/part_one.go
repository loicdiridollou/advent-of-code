package main

import (
	"math"
	"regexp"
	"strings"
)

type Rule struct {
	pattern, addition string
}

func prepareData(input string) (string, []Rule) {
	split := strings.Split(input, "\n\n")
	rules := make([]Rule, 0)

	r, _ := regexp.Compile(`(\w+) -> (\w+)`)
	for _, rule := range strings.Split(split[1], "\n") {
		if len(rule) == 0 {
			continue
		}
		rule_elements := r.FindStringSubmatch(rule)
		rules = append(rules, Rule{string(rule_elements[1]), string(rule_elements[2])})
	}

	return split[0], rules
}

func Counter(s string) map[string]int {
	dic := make(map[string]int)

	for _, el := range s {
		_, exists := dic[string(el)]
		if !exists {
			dic[string(el)] = 0
		}
		dic[string(el)] += 1
	}

	return dic
}

func copyMap(originalMap map[string]int) map[string]int {
	// Create the target map
	targetMap := make(map[string]int)

	// Copy from the original map to the target map
	for key, value := range originalMap {
		targetMap[key] = value
	}

	return targetMap
}

func getMinMax(dic map[string]int) (int, int) {
	min, max := int(math.Inf(1)), 0

	for _, j := range dic {
		if j < min {
			min = j
		}
		if j > max {
			max = j
		}
	}
	return min, max
}

func buildMap(formula string, rules []Rule) map[string]int {
	dic := make(map[string]int, 0)

	for i := 0; i < len(formula)-1; i++ {
		if _, ok := dic[formula[i:i+2]]; !ok {
			dic[formula[i:i+2]] = 0
		}
		dic[formula[i:i+2]]++
	}
	for _, rule := range rules {
		if _, ok := dic[rule.pattern]; !ok {
			dic[rule.pattern] = 0
		}
	}

	return dic
}

func part1(input string) int {
	formula, rules := prepareData(input)

	dic := buildMap(formula, rules)
	counter := Counter(formula)
	rounds := 10

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
