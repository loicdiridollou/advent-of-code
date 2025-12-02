//! # Advent of Code - Day 1 - Part Two

pub fn part2(_input: &str) -> usize {
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

    let mut dial = 50;
    let mut counter = 0;

    for num in v1.iter() {
        let dial_long = dial + num;
        counter += (dial_long / 100).abs();
        if dial != 0 && dial_long <= 0 {
            counter += 1;
        }
        dial = dial_long.rem_euclid(100);
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 6);
    }
}
