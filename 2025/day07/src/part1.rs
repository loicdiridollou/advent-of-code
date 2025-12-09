// Advent of Code 2025 - Day 06 part 1

use std::collections::HashSet;

pub fn part1(_input: &str) -> usize {
    let v1 = _input
        .lines()
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut counter = 0;
    let mut queue: Vec<(usize, usize)> = vec![];
    let mut visited: HashSet<(usize, usize)> = HashSet::new();

    for (c, val) in v1[0].iter().enumerate() {
        if *val == 'S' {
            queue.push((0, c));
            break;
        }
    }

    loop {
        if queue.is_empty() {
            break;
        }
        let curr = queue.remove(0);
        if visited.contains(&curr) || curr.0 == v1.len() {
            continue;
        } else {
            visited.insert(curr);
        }
        if v1[curr.0][curr.1] == '^' {
            queue.push((curr.0, curr.1 - 1));
            queue.push((curr.0, curr.1 + 1));
            counter += 1;
        } else {
            queue.push((curr.0 + 1, curr.1));
        }
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 4277556);
    }
}
