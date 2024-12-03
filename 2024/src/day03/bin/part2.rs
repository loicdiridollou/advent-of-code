//! # Advent of Code - Day 3 - Part Two

use regex::Regex;

pub fn part2(_input: &str) -> usize {
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();

    let mut count = 0;
    for (_, [n1, n2]) in re.captures_iter(_input).map(|c| c.extract()) {
        count += n1.parse::<usize>().unwrap() * n2.parse::<usize>().unwrap();
    }

    println!("{}", count);

    return 0;
}

#[cfg(test)]
mod day03 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 0);
    }
}
