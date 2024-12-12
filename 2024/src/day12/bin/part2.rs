//! # Advent of Code - Day 12 - Part Two

use std::collections::HashMap;
use std::collections::HashSet;

const DIRECTIONS: [[[i32; 2]; 2]; 4] = [
    [[0, 1], [1, 0]],
    [[1, 0], [0, -1]],
    [[0, -1], [-1, 0]],
    [[-1, 0], [0, 1]],
];

pub fn part2(_input: &str) -> usize {
    let input = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|item| item.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut visited = HashSet::new();
    let mut count_map = HashMap::new();
    let mut borders = HashMap::new();
    let mut points_by_group = HashMap::new();
    let mut perimeter_by_group = HashMap::new();
    let mut area_by_group = HashMap::new();
    let mut corner_by_group = HashMap::new();

    let mut group_id = 1;

    for r in 0..input.len() {
        for c in 0..input[0].len() {
            let mut count = 0;
            if visited.contains(&vec![r, c]) {
                continue;
            }
            let mut queue = vec![vec![r, c]];
            let target_char = input[r][c];
            let mut area = 0;

            while queue.len() > 0 {
                let curr = queue.remove(0);
                if visited.contains(&curr) {
                    continue;
                }
                visited.insert(curr.clone());
                borders.insert((curr[0] as i32, curr[1] as i32), group_id);
                points_by_group
                    .entry(group_id)
                    .or_insert(HashSet::new())
                    .insert((curr[0] as i32, curr[1] as i32));
                area += 1;
                let r = curr[0];
                let c = curr[1];

                let down = vec![r + 1, c];
                let right = vec![r, c + 1];
                let mut new_count = 0;

                if r > 0 {
                    let up = vec![r - 1, c];
                    if input[up[0]][up[1]] == target_char && !visited.contains(&up) {
                        queue.push(up);
                    } else if input[up[0]][up[1]] != target_char || !visited.contains(&up) {
                        new_count += 1
                    }
                } else {
                    new_count += 1;
                }
                if down[0] < input.len()
                    && input[down[0]][down[1]] == target_char
                    && !visited.contains(&down)
                {
                    queue.push(down);
                } else if down[0] == input.len() || input[down[0]][down[1]] != target_char {
                    new_count += 1;
                }
                if c > 0 {
                    let left = vec![r, c - 1];

                    if input[left[0]][left[1]] == target_char && !visited.contains(&left) {
                        queue.push(left);
                    } else if input[left[0]][left[1]] != target_char {
                        new_count += 1;
                    }
                } else if c == 0 {
                    new_count += 1;
                }
                if right[1] < input[0].len()
                    && input[right[0]][right[1]] == target_char
                    && !visited.contains(&right)
                {
                    queue.push(right);
                } else if right[1] == input[0].len() || input[right[0]][right[1]] != target_char {
                    new_count += 1;
                }
                count_map.insert(curr, new_count);
                count += new_count;
            }
            perimeter_by_group.insert(group_id, count);
            area_by_group.insert(group_id, area);
            group_id += 1;
        }
    }

    borders.iter().for_each(|(n, group_id)| {
        let count = DIRECTIONS
            .iter()
            .map(|[[x, y], [x1, y1]]| {
                let test_a = borders
                    .get(&(x + n.0, y + n.1))
                    .is_some_and(|c| c == group_id);
                let test_b = borders
                    .get(&(x1 + n.0, y1 + n.1))
                    .is_some_and(|c| c == group_id);
                if test_a
                    && test_b
                    && borders
                        .get(&(x + x1 + n.0, y + y1 + n.1))
                        .is_some_and(|c| c != group_id)
                {
                    // have interior corner
                    return 1;
                } else if !test_a && !test_b {
                    //have exterior corner
                    return 1;
                }
                return 0;
            })
            .sum::<i32>();
        *corner_by_group.entry(group_id).or_insert(0) += count;
    });

    return area_by_group
        .iter()
        .map(|(group, val)| val * corner_by_group.get(&group).unwrap())
        .sum::<i32>() as usize;
}

#[cfg(test)]
mod day12 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 1206);
    }
}
