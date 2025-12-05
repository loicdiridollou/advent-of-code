// Advent of Code 2025 - Day 04 part 1

use std::ops::Range;

pub fn part1(_input: &str) -> usize {
    let v1 = _input.split("\n\n").collect::<Vec<&str>>();
    let v2 = v1[0]
        .lines()
        .map(|c| {
            let d = c
                .split('-')
                .map(|b| b.parse::<i64>().unwrap())
                .collect::<Vec<i64>>();
            d[0]..(d[1] + 1)
        })
        .collect::<Vec<Range<i64>>>();
    let v3 = v1[1]
        .lines()
        .map(|b| b.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();

    let mut counter = 0;
    for elem in v3.iter() {
        for rng in v2.iter() {
            if rng.contains(elem) {
                counter += 1;
                break;
            }
        }
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 3);
    }
}
