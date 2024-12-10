//! # Advent of Code - Day 9 - Part One

pub fn part1(_input: &str) -> usize {
    let mut values: Vec<i64> = vec![];
    let mut id = 0;

    for (idx, el) in _input.chars().enumerate() {
        if el == '\n' {
            continue;
        } else if idx % 2 == 0 {
            for _ in 0..(el.to_string().parse::<i64>().unwrap()) {
                values.push(id);
            }
            id += 1;
        } else {
            for _ in 0..(el.to_string().parse::<i64>().unwrap()) {
                values.push(-1);
            }
        }
    }

    let mut l = 0;
    let mut r = values.len() - 1;

    while l < r {
        if values[l] != -1 {
            l += 1;
        } else if values[r] == -1 {
            r -= 1;
        } else {
            values.swap(l, r);
        }
    }
    let mut count = 0;
    for (idx, el) in values.iter().enumerate() {
        if *el == -1 {
            continue;
        }
        count += idx as i64 * el;
    }
    return count as usize;
}

#[cfg(test)]
mod day09 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 1928);
    }
}
