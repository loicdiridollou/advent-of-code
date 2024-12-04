//! # Advent of Code - Day 3 - Part One
use regex::Regex;

pub fn part1(_input: &str) -> usize {
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();

    let mut count = 0;
    for (_, [n1, n2]) in re.captures_iter(_input).map(|c| c.extract()) {
        count += n1.parse::<usize>().unwrap() * n2.parse::<usize>().unwrap();
    }

    return count;
}

#[cfg(test)]
mod day03 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))";
        assert_eq!(part1(_input), 161);
    }
}
