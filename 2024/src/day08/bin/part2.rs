//! # Advent of Code - Day 8 - Part Two

use std::collections::{HashMap, HashSet};

pub fn part2(_input: &str) -> usize {
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
        for i in 0..examples.len() {
            for j in 0..examples.len() {
                if i == j {
                    continue;
                }
                combinations.push(vec![examples[i].clone(), examples[j].clone()]);
            }
        }
        for combo in combinations.iter() {
            let dr = combo[1][0] - combo[0][0];
            let dc = combo[1][1] - combo[0][1];
            let mut r = combo[0][0];
            let mut c = combo[0][1];
            while 0 <= r && r < nr && 0 <= c && c < nc {
                echo.insert(vec![r, c]);
                r += dr;
                c += dc;
            }
        }
    }

    return echo.len();
}
#[cfg(test)]
mod day08 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 34);
    }
}
