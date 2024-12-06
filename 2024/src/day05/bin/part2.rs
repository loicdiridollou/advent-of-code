//! # Advent of Code - Day 5 - Part Two

use std::cmp::Ordering;
use std::collections::{HashMap, HashSet};

fn sort_update(update: Vec<i32>, rules_new: &HashMap<String, Ordering>) -> i32 {
    let mut upp: Vec<i32> = update.clone();
    upp.sort_by(|a, b| {
        *rules_new
            .get(&(a.to_string() + "_" + &b.to_string()))
            .unwrap_or(&Ordering::Greater)
    });

    return upp[upp.len() / 2];
}
pub fn part2(_input: &str) -> usize {
    let values = _input.split("\n\n").collect::<Vec<&str>>();
    let mut rules_map = HashMap::new();
    let mut rules_new = HashMap::new();
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

            rules_new.insert(
                val[0].to_string() + "_" + &val[1].to_string(),
                std::cmp::Ordering::Greater,
            );
            rules_new.insert(
                val[1].to_string() + "_" + &val[0].to_string(),
                std::cmp::Ordering::Less,
            );
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
    for update in updates {
        for i in 1..update.len() {
            let banned = rules_map.entry(update[i]).or_default();
            if !banned.is_empty() && update[..i].iter().any(|item| banned.contains(item)) {
                count += sort_update(update, &rules_new);
                break;
            }
        }
    }
    return count as usize;
}

#[cfg(test)]
mod day05 {
    use super::*;

    #[test]
    fn test_part2() {
        assert_eq!(part2(""), 0);
    }
}
