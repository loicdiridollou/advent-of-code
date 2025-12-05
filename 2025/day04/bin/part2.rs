//! # Advent of Code - Day 1 - Part Two

use std::collections::HashSet;

use crate::part1::Roll;

pub fn part2(_input: &str) -> usize {
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

    let mut removed = 0;
    loop {
        let mut to_remove: Vec<Roll> = vec![];
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
                removed += 1;
                to_remove.push(Roll {
                    r: roll.r,
                    c: roll.c,
                });
            }
        }
        if to_remove.is_empty() {
            break;
        } else {
            for v3 in to_remove.iter() {
                v2.remove(v3);
            }
        }
    }

    removed as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 43);
    }
}
