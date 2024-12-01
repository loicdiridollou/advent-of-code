//! # Advent of Code 202 - Day 01 part 1

pub fn part1(_input: &str) -> usize {
    let mut v1 = vec![];
    let mut v2 = vec![];

    for line in _input.lines() {
        let mut a = line.split_whitespace();
        v1.push(a.next().unwrap().parse::<i32>().unwrap());
        v2.push(a.next().unwrap().parse::<i32>().unwrap());
    }
    v1.sort();
    v2.sort();

    let result: i32 = std::iter::zip(v1, v2).map(|(l, r)| (l - r).abs()).sum();

    return result as usize;
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 11);
    }
}
