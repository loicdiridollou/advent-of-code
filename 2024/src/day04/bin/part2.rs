//! # Advent of Code - Day 4 - Part Two

use std::usize;

pub fn part2(_input: &str) -> usize {
    let maze: Vec<Vec<char>> = _input
        .split("\n")
        .filter(|&x| !x.is_empty())
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect();

    let mut _count = 0;
    for i in 1..(maze.len() - 1) {
        for j in 1..(maze[0].len() - 1) {
            if maze[i][j] == 'A' {
                if validate_pattern(&maze, i as i32, j as i32) {
                    _count += 1;
                }
            }
        }
    }

    return _count as usize;
}

fn validate_pattern(maze: &Vec<Vec<char>>, r: i32, c: i32) -> bool {
    if maze[(r - 1) as usize][(c - 1) as usize] == 'M'
        && maze[(r - 1) as usize][(c + 1) as usize] == 'S'
        && maze[(r + 1) as usize][(c - 1) as usize] == 'M'
        && maze[(r + 1) as usize][(c + 1) as usize] == 'S'
    {
        return true;
    }
    if maze[(r - 1) as usize][(c - 1) as usize] == 'M'
        && maze[(r - 1) as usize][(c + 1) as usize] == 'M'
        && maze[(r + 1) as usize][(c - 1) as usize] == 'S'
        && maze[(r + 1) as usize][(c + 1) as usize] == 'S'
    {
        return true;
    }
    if maze[(r - 1) as usize][(c - 1) as usize] == 'S'
        && maze[(r - 1) as usize][(c + 1) as usize] == 'M'
        && maze[(r + 1) as usize][(c - 1) as usize] == 'S'
        && maze[(r + 1) as usize][(c + 1) as usize] == 'M'
    {
        return true;
    }
    if maze[(r - 1) as usize][(c - 1) as usize] == 'S'
        && maze[(r - 1) as usize][(c + 1) as usize] == 'S'
        && maze[(r + 1) as usize][(c - 1) as usize] == 'M'
        && maze[(r + 1) as usize][(c + 1) as usize] == 'M'
    {
        return true;
    }
    return false;
}

#[cfg(test)]
mod day04 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 9);
    }
}
