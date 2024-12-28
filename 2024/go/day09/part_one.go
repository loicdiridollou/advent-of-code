package main

import (
	"strconv"
)

func parse_string(input string) []string {
	new_string := []string{}
	current := "0"
	for i := 0; i < len(input); i += 2 {
		value, _ := strconv.Atoi(string(input[i]))
		idx := 0
		for idx < value {
			new_string = append(new_string, current)
			idx += 1
		}
		if i+1 < len(input) {
			value, _ := strconv.Atoi(string(input[i+1]))
			idx := 0
			for idx < value {
				new_string = append(new_string, ".")
				idx += 1
			}
		}
		value, _ = strconv.Atoi(current)
		current = strconv.Itoa(value + 1)
	}

	return new_string
}

// part1 function
func part1(input string) int {
	parsed_str := parse_string(input)

	i, j := 0, len(parsed_str)-1

	for i < j {
		if parsed_str[i] != "." {
			i += 1
		} else if parsed_str[j] == "." {
			j -= 1
		} else {
			parsed_str[i], parsed_str[j] = parsed_str[j], parsed_str[i]
		}
	}

	count := 0
	for i := 0; i < len(parsed_str); i++ {
		if parsed_str[i] == "." {
			break
		}
		value, _ := strconv.Atoi(string(parsed_str[i]))
		count += i * value
	}

	return count
}
