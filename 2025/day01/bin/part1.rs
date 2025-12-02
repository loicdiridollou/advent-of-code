//! # Advent of Code 202 - Day 01 part 1

pub fn part1(_input: &str) -> usize {
    let mut v1: Vec<i32> = vec![];
    _input.lines().for_each(|line| {
        let a = line.chars().next().unwrap();
        let b = &line[1..].parse::<i32>().unwrap();
        if a == 'R' {
            v1.push(*b);
        } else {
            v1.push(-b);
        };
    });

    let mut result = 50;
    let mut counter = 0;

    for val in v1.iter() {
        result = (result + val) % 100;
        if result < 0 {
            result += 100;
        } else if result == 0 {
            counter += 1;
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
