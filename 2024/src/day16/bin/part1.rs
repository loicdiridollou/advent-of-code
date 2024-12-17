//! # Advent of Code - Day 16 - Part One

use std::{
    cmp::Ordering,
    collections::{BinaryHeap, HashMap, HashSet},
};

#[derive(Copy, Clone, Eq, PartialEq)]
struct State {
    cost: usize,
    position: (usize, usize),
    direction: (i32, i32),
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        // Notice that we flip the ordering on costs.
        // In case of a tie we compare positions - this step is necessary
        // to make implementations of `PartialEq` and `Ord` consistent.
        other
            .cost
            .cmp(&self.cost)
            .then_with(|| self.position.cmp(&other.position))
    }
}

// `PartialOrd` needs to be implemented as well.
impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

pub fn part1(_input: &str) -> usize {
    let grid = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    // for val in grid.iter() {
    //     println!("{:?}", val);
    // }
    let mut mapping = HashMap::new();
    let dir = (0, 1);
    let mut start = (0, 0);
    let mut end = (0, 0);

    for r in 0..grid.len() {
        for c in 0..grid[0].len() {
            let mut chr = grid[r][c];
            if chr == 'S' {
                start = (r, c);
                chr = '.';
            } else if chr == 'E' {
                end = (r, c);
            }
            mapping.insert((r, c), chr);
        }
    }
    let mut visited = HashSet::new();

    let mut heap = BinaryHeap::new();

    heap.push(State {
        cost: 0,
        position: start,
        direction: dir,
    });

    while let Some(State {
        cost,
        position,
        direction,
    }) = heap.pop()
    {
        if position == end {
            return Some(cost).unwrap();
        }
        visited.insert((position, direction));

        for potential in [
            State {
                cost: cost + 1,
                position: (
                    (position.0 as i32 + direction.0) as usize,
                    (position.1 as i32 + direction.1) as usize,
                ),
                direction,
            },
            State {
                cost: cost + 1000,
                position,
                direction: (direction.1, -direction.0),
            },
            State {
                cost: cost + 1000,
                position,
                direction: (-direction.1, direction.0),
            },
        ] {
            let next = State {
                cost: potential.cost,
                position: potential.position,
                direction: potential.direction,
            };

            if *mapping.get(&potential.position).unwrap() == '#' {
                continue;
            }
            if visited.contains(&(potential.position, potential.direction)) {
                continue;
            }

            heap.push(next);
        }
    }

    return 0;
}

#[cfg(test)]
mod day16 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 7036);
    }
}
