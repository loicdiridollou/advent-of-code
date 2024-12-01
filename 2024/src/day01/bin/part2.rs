//! # Advent of Code - Day 1 - Part Two

use counter::Counter;

pub fn part2(_input: &str) -> usize {
    let mut v1 = Vec::<i32>::new();
    let mut v2 = Vec::<i32>::new();
    for line in _input.lines() {
        let a: Vec<&str> = line.split_whitespace().collect();
        v1.push(a[0].parse().unwrap());
        v2.push(a[1].parse().unwrap());
    }

    let count = v2.iter().collect::<Counter<_>>();

    let mut diff = 0;
    for ai in v1.iter() {
        diff += ai * (count[&ai] as i32);
    }

    return diff as usize;
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(&_input), 31);
    }
}
