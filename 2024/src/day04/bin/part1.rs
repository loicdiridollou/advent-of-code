//! # Advent of Code - Day 4 - Part One

use std::collections::HashMap;

pub fn part1(_input: &str) -> usize {
    let maze: Vec<Vec<char>> = _input
        .split("\n")
        .filter(|&x| !x.is_empty())
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect();

    let directions = HashMap::from([
        ("N", (-1, 0)),
        ("S", (1, 0)),
        ("E", (0, 1)),
        ("W", (0, -1)),
        ("NW", (-1, -1)),
        ("NE", (-1, 1)),
        ("SW", (1, -1)),
        ("SE", (1, 1)),
    ]);

    let mut count = 0;
    for i in 0..maze.len() {
        for j in 0..maze[0].len() {
            if maze[i][j] == 'X' {
                for (_, (dr, dc)) in &directions {
                    if explore(&maze, i as i32, j as i32, *dr as i32, *dc as i32, "XMAS") {
                        count += 1;
                    }
                }
            }
        }
    }

    return count as usize;
}

fn explore(maze: &Vec<Vec<char>>, r: i32, c: i32, dr: i32, dc: i32, chr: &str) -> bool {
    let expected = chr.chars().next().unwrap();
    if maze[r as usize][c as usize] == expected && chr == "S" {
        // we have found the sequence
        return true;
    }
    if maze[r as usize][c as usize] != expected {
        // wrong character, exit early
        return false;
    }
    if r + dr < maze.len() as i32 && r + dr >= 0 && c + dc < maze[0].len() as i32 && c + dc >= 0 {
        return explore(maze, r + dr, c + dc, dr, dc, &chr[1..].to_string());
    }

    return false;
}

#[cfg(test)]
mod day04 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 18);
    }
}
