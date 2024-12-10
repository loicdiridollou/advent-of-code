//! # Advent of Code - Day 10 - Part Two

fn explore_maze(maze: &Vec<Vec<u32>>, r: usize, c: usize) -> usize {
    let expected = maze[r][c] + 1;
    if maze[r][c] == 9 {
        return 1;
    }
    let mut count = 0;

    if r > 0 && maze[r - 1][c] == expected {
        count += explore_maze(maze, r - 1, c);
    }
    if r < maze.len() - 1 && maze[r + 1][c] == expected {
        count += explore_maze(maze, r + 1, c);
    }
    if c > 0 && maze[r][c - 1] == expected {
        count += explore_maze(maze, r, c - 1);
    }
    if c < maze[0].len() - 1 && maze[r][c + 1] == expected {
        count += explore_maze(maze, r, c + 1);
    }

    return count;
}

pub fn part2(_input: &str) -> usize {
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
                count += explore_maze(&maze, r, c);
            }
        }
    }
    return count;
}

#[cfg(test)]
mod day10 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 81);
    }
}
