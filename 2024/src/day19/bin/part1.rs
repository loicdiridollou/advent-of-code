//! # Advent of Code - Day 19 - Part One

use std::collections::HashSet;

use cached::proc_macro::cached;

#[cached]
fn is_doable(towel: String, patterns_str: String, max_size: usize) -> bool {
    if towel == "" {
        return true;
    }
    let mut patterns = HashSet::new();
    let patternss = patterns_str.clone();
    patternss.split(", ").for_each(|c| {
        patterns.insert(c);
    });

    for i in 1..=max_size {
        if i > towel.len() {
            return false;
        }
        if patterns.contains(&towel[..i])
            && is_doable(towel[i..].to_string(), patterns_str.clone(), max_size)
        {
            return true;
        }
    }
    return false;
}

pub fn part1(_input: &str) -> usize {
    let input = _input.split("\n\n").collect::<Vec<&str>>();
    let mut patterns = HashSet::new();

    let max_size = input[0]
        .split(", ")
        .map(|c| {
            patterns.insert(c);
            return c.len();
        })
        .max()
        .unwrap();

    let towels = input[1]
        .split("\n")
        .filter(|c| !c.is_empty())
        .collect::<Vec<&str>>();

    let mut count = 0;
    for towel in towels {
        if is_doable(towel.to_string(), input[0].to_string(), max_size) {
            count += 1;
        }
    }
    return count;
}

#[cfg(test)]
mod day19 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 6);
    }
}
