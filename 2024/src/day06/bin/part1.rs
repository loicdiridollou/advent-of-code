//! # Advent of Code - Day 6 - Part One

use std::collections::{HashMap, HashSet};

pub fn move_guard(
    mut row: i32,
    mut col: i32,
    nr: i32,
    nc: i32,
    mut direction: &str,
    obstacles: HashSet<Vec<i32>>,
) -> HashSet<Vec<i32>> {
    let mut visited: HashSet<Vec<i32>> = HashSet::new();
    let dir_map = HashMap::from([("U", [-1, 0]), ("D", [1, 0]), ("L", [0, -1]), ("R", [0, 1])]);
    let directions = vec!["U", "R", "D", "L"];

    visited.insert(vec![row, col]);

    loop {
        let dpos = dir_map.get(direction).unwrap();
        let newr = row + dpos[0];
        let newc = col + dpos[1];

        if (newr < 0) || (newr >= nr) || (newc < 0) || (newc >= nc) {
            break;
        } else if obstacles.contains(&vec![newr, newc]) {
            direction =
                directions[(directions.iter().position(|&x| x == direction).unwrap() + 1) % 4];
        } else {
            row = newr;
            col = newc;
        }

        visited.insert(vec![row, col]);
    }

    return visited;
}

pub fn part1(_input: &str) -> usize {
    let maze: Vec<Vec<char>> = _input
        .split("\n")
        .filter(|&x| !x.is_empty())
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect();

    let mut obstacles = HashSet::new();

    let mut r: usize = 42;
    let mut c: usize = 42;
    for i in 0..maze.len() {
        for j in 0..maze[0].len() {
            if maze[i][j] == '#' {
                obstacles.insert(vec![i as i32, j as i32]);
            } else if maze[i][j] == '^' {
                r = i;
                c = j;
            }
        }
    }

    return move_guard(
        r as i32,
        c as i32,
        maze.len() as i32,
        maze[0].len() as i32,
        "U",
        obstacles,
    )
    .len();
}

#[cfg(test)]
mod day06 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 41);
    }
}
