//! # Advent of Code - Day 8 - Part One

use std::collections::{HashMap, HashSet};

fn is_in_maze(point: &Vec<i32>, nr: i32, nc: i32) -> bool {
    return point[0] >= 0 && point[0] < nr && point[1] >= 0 && point[1] < nc;
}

pub fn part1(_input: &str) -> usize {
    let maze: Vec<Vec<char>> = _input
        .split("\n")
        .filter(|&x| !x.is_empty())
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect();

    let mut antennas = HashMap::new();
    let nr = maze.len() as i32;
    let nc = maze[0].len() as i32;
    for r in 0..maze.len() {
        for c in 0..maze[0].len() {
            if maze[r][c] != '.' {
                antennas
                    .entry(maze[r][c])
                    .or_insert(vec![])
                    .push(vec![r as i32, c as i32]);
            }
        }
    }
    let mut echo = HashSet::new();
    for antenna in antennas.keys() {
        let examples = antennas.get(&antenna).unwrap().to_vec();
        let mut combinations = vec![];
        for i in 0..(examples.len() - 1) {
            for j in (i + 1)..examples.len() {
                combinations.push(vec![examples[i].clone(), examples[j].clone()]);
            }
        }
        for combo in combinations.iter() {
            let first = vec![2 * combo[0][0] - combo[1][0], 2 * combo[0][1] - combo[1][1]];
            let second = vec![2 * combo[1][0] - combo[0][0], 2 * combo[1][1] - combo[0][1]];

            if is_in_maze(&first, nr, nc) {
                echo.insert(first.clone());
            }
            if is_in_maze(&second, nr, nc) {
                echo.insert(second.clone());
            }
        }
    }

    return echo.len();
}

#[cfg(test)]
mod day08 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 14);
    }
}
