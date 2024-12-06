//! # Advent of Code - Day 5 - Part One

use std::collections::{HashMap, HashSet};

pub fn part1(_input: &str) -> usize {
    let values = _input.split("\n\n").collect::<Vec<&str>>();
    let mut rules_map = HashMap::new();
    values[0]
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|c| {
            c.split("|")
                .map(|x| x.parse::<i32>().unwrap())
                .collect::<Vec<i32>>()
        })
        .for_each(|val| {
            rules_map
                .entry(val[0])
                .or_insert(HashSet::new())
                .insert(val[1]);
        });

    let updates = values[1]
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|c| {
            c.split(",")
                .map(|x| x.parse::<i32>().unwrap())
                .collect::<Vec<i32>>()
        })
        .collect::<Vec<Vec<i32>>>();

    let mut count = 0;
    updates.iter().for_each(|update| {
        let mut valid = true;
        for i in 1..update.len() {
            let banned = rules_map.entry(update[i]).or_default();
            if !banned.is_empty() && update[..i].iter().any(|item| banned.contains(item)) {
                valid = false;
                break;
            }
        }
        if valid {
            count += update[update.len() / 2]
        }
    });
    return count as usize;
}

#[cfg(test)]
mod day05 {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(""), 0);
    }
}
