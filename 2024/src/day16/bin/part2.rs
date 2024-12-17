//! # Advent of Code - Day 16 - Part Two

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

pub fn part2(_input: &str) -> usize {
    let grid = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

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
    let mut backtrack = HashMap::new();
    let mut lowest_cost: HashMap<((usize, usize), (i32, i32)), usize> = HashMap::new();
    let mut end_states = HashSet::new();
    let mut best_cost = 100000000000;

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
            if Some(cost).unwrap() > best_cost {
                break;
            }
            best_cost = Some(cost).unwrap();
            end_states.insert((position, dir));
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
            let lowest = *lowest_cost
                .get(&(potential.position, potential.direction))
                .unwrap_or(&100000000000);

            if potential.cost > lowest {
                continue;
            }
            let entry = (potential.position, potential.direction);
            if potential.cost < lowest {
                *lowest_cost
                    .entry((potential.position, potential.direction))
                    .or_insert(0) = potential.cost;
                *backtrack.entry(entry).or_insert(HashSet::new()) = HashSet::new();
            }
            backtrack
                .entry((potential.position, potential.direction))
                .or_insert(HashSet::new())
                .insert((position, direction));

            heap.push(next);
        }
    }

    let mut states = vec![];
    let mut seen = HashSet::new();
    for el in end_states.iter() {
        states.push(el);
        seen.insert(el);
    }
    while states.len() > 0 {
        let key = states.remove(0);
        for last in backtrack.get(key).unwrap() {
            if seen.contains(&last) {
                continue;
            }
            seen.insert(last);
            states.push(last);
        }
    }

    let mut visited = HashSet::new();
    for el in seen.iter() {
        visited.insert(el.0);
    }

    return visited.len() as usize;
}

#[cfg(test)]
mod day16 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 45);
    }
}
