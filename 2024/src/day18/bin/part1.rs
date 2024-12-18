//! # Advent of Code - Day 18 - Part One

use std::collections::HashMap;
use std::collections::HashSet;

pub fn part1(_input: &str, parameters: (i32, i32, (i32, i32)), num_bits: usize) -> usize {
    let mut mapping = HashMap::new();
    let _input = _input
        .split("\n")
        .filter(|c| !c.is_empty())
        .map(|v| {
            let values = v.split(",").collect::<Vec<&str>>();
            return (
                values[1].parse::<i32>().unwrap(),
                values[0].parse::<i32>().unwrap(),
            );
        })
        .collect::<Vec<(i32, i32)>>();
    for i in 0.._input.len() {
        mapping.insert(_input[i], i);
    }

    let (_, min_steps) = run_program(&mapping, parameters, num_bits);

    return min_steps;
}

pub fn run_program(
    mapping: &HashMap<(i32, i32), usize>,
    parameters: (i32, i32, (i32, i32)),
    limit: usize,
) -> (bool, usize) {
    let mut queue = vec![(0, 0, 0)];
    let mut visited = HashSet::new();
    visited.insert((0, 0));
    let numr = parameters.0;
    let numc = parameters.1;
    let end = parameters.2;
    let mut found = false;
    let mut value = 0;
    while queue.len() > 0 {
        let curr = queue.remove(0);
        for (dr, dc) in [(-1, 0), (0, 1), (1, 0), (0, -1)] {
            let nr = curr.0 + dr;
            let nc = curr.1 + dc;
            if nr < 0 || nc < 0 || nr >= numr || nc >= numc {
                continue;
            }
            if visited.contains(&(nr, nc)) {
                continue;
            }

            if (nr, nc) == end {
                value = curr.2 + 1;
                found = true;
                break;
            }
            if *mapping.get(&(curr.0, curr.1)).unwrap_or(&(limit + 2)) < limit {
                continue;
            }
            queue.push((nr, nc, curr.2 + 1));
            visited.insert((nr, nc));
        }
        if found {
            break;
        }
    }
    return (found, value);
}

#[cfg(test)]
mod day18 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = "5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0";
        assert_eq!(part1(_input, (7, 7, (6, 6)), 12), 22);
    }
}
