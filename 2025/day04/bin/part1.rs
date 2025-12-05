// Advent of Code 2025 - Day 04 part 1

use std::collections::HashSet;

#[derive(Eq, PartialEq, Debug, Hash)]
pub struct Roll {
    pub r: i32,
    pub c: i32,
}

pub fn part1(_input: &str) -> usize {
    let v1 = _input
        .lines()
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut v2: HashSet<Roll> = HashSet::new();
    for (i, line) in v1.iter().enumerate() {
        for (j, item) in line.iter().enumerate() {
            if *item == '@' {
                v2.insert(Roll {
                    r: i as i32,
                    c: j as i32,
                });
            }
        }
    }

    let directions = [
        (0, 1),
        (1, 0),
        (1, 1),
        (0, -1),
        (-1, 0),
        (-1, -1),
        (1, -1),
        (-1, 1),
    ];

    let mut counter = 0;
    for roll in v2.iter() {
        let mut dir_check = 0;
        for dir in directions.iter() {
            if v2.contains(&Roll {
                r: (roll.r + dir.0),
                c: (roll.c + dir.1),
            }) {
                dir_check += 1;
            }
        }
        if dir_check < 4 {
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
        assert_eq!(part1(_input), 13);
    }
}
