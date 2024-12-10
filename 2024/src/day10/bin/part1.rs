//! # Advent of Code - Day 10 - Part One

use std::collections::HashSet;

fn explore_maze(maze: &Vec<Vec<u32>>, r: usize, c: usize) -> HashSet<Vec<usize>> {
    let expected = maze[r][c] + 1;
    let mut hash_set = HashSet::new();
    if maze[r][c] == 9 {
        hash_set.insert(vec![r, c]);
        return hash_set;
    }

    if r > 0 && maze[r - 1][c] == expected {
        hash_set.extend(explore_maze(maze, r - 1, c));
    }
    if r < maze.len() - 1 && maze[r + 1][c] == expected {
        hash_set.extend(explore_maze(maze, r + 1, c));
    }
    if c > 0 && maze[r][c - 1] == expected {
        hash_set.extend(explore_maze(maze, r, c - 1));
    }
    if c < maze[0].len() - 1 && maze[r][c + 1] == expected {
        hash_set.extend(explore_maze(maze, r, c + 1));
    }

    return hash_set;
}

pub fn part1(_input: &str) -> usize {
    let maze = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|item| {
            item.chars()
                .map(|c| c.to_digit(10).unwrap())
                .collect::<Vec<u32>>()
        })
        .collect::<Vec<Vec<u32>>>();

    let mut count = 0;

    for r in 0..maze.len() {
        for c in 0..maze[0].len() {
            if maze[r][c] == 0 {
                count += explore_maze(&maze, r, c).len();
            }
        }
    }
    return count;
}

#[cfg(test)]
mod day10 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 36);
    }
}
