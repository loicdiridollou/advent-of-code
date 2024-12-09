//! # Advent of Code - Day 6 - Part Two

use std::collections::{HashMap, HashSet};

use crate::part1::move_guard;

#[derive(Eq, Hash, PartialEq, Debug)]
struct Position {
    row: i32,
    col: i32,
    dir: String,
}

fn identify_loop(
    mut row: i32,
    mut col: i32,
    nr: i32,
    nc: i32,
    mut direction: &str,
    obstacles: &HashSet<Vec<i32>>,
) -> bool {
    let mut visited: HashSet<Position> = HashSet::new();
    let dir_map = HashMap::from([("U", [-1, 0]), ("D", [1, 0]), ("L", [0, -1]), ("R", [0, 1])]);
    let directions = vec!["U", "R", "D", "L"];

    visited.insert(Position {
        row,
        col,
        dir: direction.to_string(),
    });
    let mut is_loop: bool = false;

    loop {
        let dpos = dir_map.get(direction).unwrap();
        let newr = row + dpos[0];
        let newc = col + dpos[1];

        if (newr < 0) || (newr >= nr) || (newc < 0) || (newc >= nc) {
            break;
        } else if obstacles.contains(&vec![newr, newc]) {
            direction =
                directions[(directions.iter().position(|&x| x == direction).unwrap() + 1) % 4];
        } else if visited.contains(&Position {
            row: newr,
            col: newc,
            dir: direction.to_string(),
        }) {
            is_loop = true;
            break;
        } else {
            row = newr;
            col = newc;
        }

        visited.insert(Position {
            row,
            col,
            dir: direction.to_string(),
        });
    }

    return is_loop;
}

pub fn part2(_input: &str) -> usize {
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

    let early_visits = move_guard(
        r as i32,
        c as i32,
        maze.len() as i32,
        maze[0].len() as i32,
        "U",
        obstacles.clone(),
    );
    let mut possible_positions = HashSet::new();

    for el in early_visits.iter() {
        obstacles.insert(el.to_vec());
        if identify_loop(
            r as i32,
            c as i32,
            maze.len() as i32,
            maze[0].len() as i32,
            "U",
            &obstacles.clone(),
        ) {
            possible_positions.insert(el);
        }
        obstacles.remove(el);
    }

    return possible_positions.len();
}

#[cfg(test)]
mod day06 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 6);
    }
}
