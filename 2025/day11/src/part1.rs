// Advent of Code 2025 - Day 06 part 1

use std::collections::{HashMap, HashSet};

pub fn part1(_input: &str) -> usize {
    let mut mapping: HashMap<&str, Vec<&str>> = HashMap::new();

    _input.lines().for_each(|c| {
        let v = c.split(':').collect::<Vec<&str>>();
        let elem = v[1].split_whitespace().collect::<Vec<&str>>();
        mapping.insert(v[0], elem);
    });

    let mut queue = vec!["you"];
    let mut visited: HashSet<&str> = HashSet::new();
    let mut counter = 0;

    while !queue.is_empty() {
        let curr = queue.remove(0);
        visited.insert(curr);

        for dest in mapping.get(&curr).unwrap().iter() {
            if *dest == "out" {
                counter += 1;
                continue;
            }
            queue.push(dest);
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
