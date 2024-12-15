//! # Advent of Code - Day 15 - Part Two

use std::collections::{HashMap, HashSet};

use crate::part1::convert_rule_to_direction;

pub fn part2(_input: &str) -> usize {
    let mut mapping = HashMap::new();
    let input = _input.split("\n\n").collect::<Vec<&str>>();

    let map_str = input[0]
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    let mut curr = (-1, -1);

    for i in 0..map_str.len() {
        for j in 0..map_str[0].len() {
            if map_str[i][j] == '@' {
                curr = (i as i32, 2 * j as i32);
                mapping.insert((i as i32, 2 * j as i32), '.');
                mapping.insert((i as i32, (2 * j + 1) as i32), '.');
            } else if map_str[i][j] == '#' {
                mapping.insert((i as i32, 2 * j as i32), '#');
                mapping.insert((i as i32, (2 * j + 1) as i32), '#');
            } else if map_str[i][j] == '.' {
                mapping.insert((i as i32, 2 * j as i32), '.');
                mapping.insert((i as i32, (2 * j + 1) as i32), '.');
            } else if map_str[i][j] == 'O' {
                mapping.insert((i as i32, 2 * j as i32), '[');
                mapping.insert((i as i32, (2 * j + 1) as i32), ']');
            }
        }
    }

    let rules = input[1].replace("\n", "").chars().collect::<Vec<char>>();

    for rule in rules.iter() {
        let dir = convert_rule_to_direction(&rule);
        let (dr, dc) = dir;
        let mut targets = vec![curr];
        let mut visited = HashSet::new();
        let mut movable = true;
        let mut idx = 0;

        while idx < targets.len() {
            let (r, c) = targets[idx];
            let nr = r + dr;
            let nc = c + dc;

            if visited.contains(&(nr, nc)) {
                idx += 1;
                continue;
            }
            visited.insert((nr, nc));
            let chr = *mapping.get(&(nr, nc)).unwrap();
            if chr == '#' {
                movable = false;
                break;
            }
            if chr == ']' {
                targets.push((nr, nc));
                targets.push((nr, nc - 1));
            }
            if chr == '[' {
                targets.push((nr, nc));
                targets.push((nr, nc + 1));
            }
            idx += 1
        }
        if !movable {
            continue;
        }
        let copy = mapping.clone();
        for (br, bc) in targets.iter() {
            *mapping.get_mut(&(*br, *bc)).unwrap() = '.';
        }
        curr = (curr.0 + dr, curr.1 + dc);
        for (br, bc) in targets[1..].iter() {
            *mapping.get_mut(&(*br + dr, *bc + dc)).unwrap() = *copy.get(&(*br, *bc)).unwrap();
        }
        // display_grid(&mapping, curr, map_str.len(), 2 * map_str[0].len());
    }

    let total = mapping
        .iter()
        .map(|((r, c), chr)| if *chr == '[' { r * 100 + c } else { 0 })
        .sum::<i32>();

    return total as usize;
}

#[cfg(test)]
mod day15 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 9021);
    }
}
