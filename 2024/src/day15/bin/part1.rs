//! # Advent of Code - Day 15 - Part One

use std::collections::HashMap;

pub fn convert_rule_to_direction(chr: &char) -> (i32, i32) {
    match chr {
        '<' => (0, -1),
        '^' => (-1, 0),
        'v' => (1, 0),
        '>' => (0, 1),
        _ => (0, 0),
    }
}

#[allow(dead_code)]
pub fn display_grid(mapping: &HashMap<(i32, i32), char>, curr: (i32, i32), nr: usize, nc: usize) {
    for i in 0..(nr as i32) {
        let mut grid = vec![];
        for j in 0..(nc as i32) {
            if (i, j) == curr {
                grid.push('@');
            } else {
                grid.push(*mapping.get(&(i, j)).unwrap())
            }
        }
        println!("{:?}", grid);
    }
}

pub fn part1(_input: &str) -> usize {
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
                curr = (i as i32, j as i32);
                mapping.insert((i as i32, j as i32), '.');
            } else {
                mapping.insert((i as i32, j as i32), map_str[i][j]);
            }
        }
    }

    let rules = input[1].replace("\n", "").chars().collect::<Vec<char>>();

    for rule in rules.iter() {
        let dir = convert_rule_to_direction(&rule);
        (curr, mapping) = apply_rule(curr, dir, mapping);
    }

    let total = mapping
        .iter()
        .map(|((r, c), chr)| if *chr == 'O' { r * 100 + c } else { 0 })
        .sum::<i32>();

    return total as usize;
}

fn apply_rule(
    curr: (i32, i32),
    dir: (i32, i32),
    mut mapping: HashMap<(i32, i32), char>,
) -> ((i32, i32), HashMap<(i32, i32), char>) {
    let new_pos = (curr.0 + dir.0, curr.1 + dir.1);
    let mut current = curr;
    if *mapping.get(&new_pos).unwrap() == '#' {
        // it is blocked we can't do anything
    } else if *mapping.get(&new_pos).unwrap() == '.' {
        // it is free we can move
        current = new_pos;
    } else if *mapping.get(&new_pos).unwrap() == 'O' {
        // there is something so we try to move it
        // we try to find space before the wall
        let mut pos = new_pos;
        let mut movable = false;
        loop {
            if *mapping.get(&pos).unwrap() == 'O' {
                pos = (pos.0 + dir.0, pos.1 + dir.1);
            } else if *mapping.get(&pos).unwrap() == '.' {
                movable = true;
                break;
            } else {
                // everything is already stacked
                break;
            }
        }
        if movable {
            *mapping.get_mut(&new_pos).unwrap() = '.';
            *mapping.get_mut(&pos).unwrap() = 'O';
            current = new_pos;
        }
    }
    return (current, mapping);
}

#[cfg(test)]
mod day15 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 10092);
    }
}
