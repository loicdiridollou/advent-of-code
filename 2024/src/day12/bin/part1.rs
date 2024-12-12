//! # Advent of Code - Day 12 - Part One

use std::collections::HashMap;
use std::collections::HashSet;

pub fn part1(_input: &str) -> usize {
    let input = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|item| item.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut visited = HashSet::new();
    let mut perimeter_by_group = HashMap::new();
    let mut area_by_group = HashMap::new();

    let mut group_id = 1;
    let nr = input.len();
    let nc = input[0].len();

    for r in 0..input.len() {
        for c in 0..input[0].len() {
            if visited.contains(&(r, c)) {
                continue;
            }
            let mut new_count = 0;
            let mut queue = vec![(r, c)];
            let target_char = input[r][c];
            let mut area = 0;

            while queue.len() > 0 {
                let curr = queue.remove(0);
                if visited.contains(&curr) {
                    continue;
                }
                visited.insert(curr.clone());
                area += 1;
                let (r, c) = curr;

                let down = (r + 1, c);
                let right = (r, c + 1);

                // up
                if r > 0 {
                    let up = (r - 1, c);
                    if input[up.0][up.1] == target_char && !visited.contains(&up) {
                        queue.push(up);
                    } else if input[up.0][up.1] != target_char || !visited.contains(&up) {
                        new_count += 1
                    }
                } else {
                    new_count += 1;
                }

                // down
                if down.0 < nr && input[down.0][down.1] == target_char && !visited.contains(&down) {
                    queue.push(down);
                } else if down.0 == input.len() || input[down.0][down.1] != target_char {
                    new_count += 1;
                }

                // left
                if c > 0 {
                    let left = (r, c - 1);

                    if input[left.0][left.1] == target_char && !visited.contains(&left) {
                        queue.push(left);
                    } else if input[left.0][left.1] != target_char {
                        new_count += 1;
                    }
                } else if c == 0 {
                    new_count += 1;
                }

                // right
                if right.1 < nc
                    && input[right.0][right.1] == target_char
                    && !visited.contains(&right)
                {
                    queue.push(right);
                } else if right.1 == input[0].len() || input[right.0][right.1] != target_char {
                    new_count += 1;
                }
            }
            perimeter_by_group.insert(group_id, new_count);
            area_by_group.insert(group_id, area);
            group_id += 1;
        }
    }

    let res = area_by_group
        .iter()
        .map(|(group, val)| val * perimeter_by_group.get(&group).unwrap())
        .sum::<i32>();

    return res as usize;
}

#[cfg(test)]
mod day12 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 1930);
    }
}
