//! # Advent of Code - Day 25 - Part One

use itertools::Itertools;

pub fn part1(_input: &str) -> usize {
    let entries = _input.split("\n\n").collect::<Vec<&str>>();
    let mut keys = vec![];
    let mut locks = vec![];

    for entry in entries {
        let grid = entry
            .split("\n")
            .filter(|c| !c.is_empty())
            .map(|c| c.chars().collect::<Vec<char>>())
            .collect::<Vec<Vec<char>>>();
        let mut res = vec![];
        if grid[0].iter().all(|c| *c == '#') {
            while res.len() < grid[0].len() {
                for i in 1..=grid.len() {
                    if grid[grid.len() - i][res.len()] == '#' {
                        res.push(grid.len() - i);
                        break;
                    }
                }
            }
            locks.push(res);
        } else {
            while res.len() < grid[0].len() {
                for i in 0..grid.len() {
                    if grid[i][res.len()] == '#' {
                        res.push(grid.len() - i - 1);
                        break;
                    }
                }
            }
            keys.push(res);
        }
    }

    let count = keys
        .iter()
        .cartesian_product(locks.iter())
        .map(|(key, lock)| {
            if lock.iter().zip(key.iter()).all(|(x, y)| x + y <= 5) {
                return 1;
            } else {
                return 0;
            }
        })
        .sum();

    return count;
}

#[cfg(test)]
mod day25 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 3);
    }
}
