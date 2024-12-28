package main

import (
	"slices"
	"strconv"
)

func parse_string_new(input string) ([][2]int, int) {
	fs_arr := [][2]int{}
	current := 0
	for i := 0; i < len(input); i += 2 {
		value, _ := strconv.Atoi(string(input[i]))
		fs_arr = append(fs_arr, [2]int{current, value})
		if i+1 < len(input) {
			value, _ := strconv.Atoi(string(input[i+1]))
			fs_arr = append(fs_arr, [2]int{-1, value})
		}
		current += 1
	}

	return fs_arr, current - 1
}

func clean_up_string(arr [][2]int) [][2]int {
	idx := 0
	for idx < len(arr)-1 {
		if arr[idx][0] != -1 {
			idx += 1
		} else if arr[idx][0] == -1 && idx < len(arr)-1 && arr[idx+1][0] == -1 {
			arr[idx] = [2]int{arr[idx][0], arr[idx][1] + arr[idx+1][1]}
			arr = append(arr[:idx+1], arr[idx+2:]...)
		} else {
			idx += 1
		}
	}
	return arr
}

func move_block(parsed_str [][2]int, id int) [][2]int {
	i, j := 0, len(parsed_str)-1

	for i < j {
		if parsed_str[i][0] != -1 {
			i += 1
		} else if parsed_str[j][0] != id {
			j -= 1
		} else if parsed_str[j][0] == id && parsed_str[i][1] >= parsed_str[j][1] {
			// do the changes
			remaining := parsed_str[i][1] - parsed_str[j][1]
			parsed_str[i] = parsed_str[j]
			parsed_str[j] = [2]int{-1, parsed_str[j][1]}

			if remaining > 0 {
				parsed_str = slices.Insert(parsed_str, i+1, [2]int{-1, remaining})
			}
			return clean_up_string(parsed_str)
		} else {
			i += 1
		}
	}
	return parsed_str
}

// part2 function
func part2(input string) int {
	parsed_str, max_id := parse_string_new(input)

	for i := max_id; i >= 0; i-- {
		parsed_str = move_block(parsed_str, i)
	}

	count := 0
	idx := 0
	for i := 0; i < len(parsed_str); i++ {
		if parsed_str[i][0] != -1 {
			for val := 1; val <= parsed_str[i][1]; val++ {
				count += idx * parsed_str[i][0]
				idx += 1
			}
		} else {
			idx += parsed_str[i][1]
		}
	}

	return count
}
