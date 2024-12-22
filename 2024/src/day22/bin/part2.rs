//! # Advent of Code - Day 22 - Part Two

use std::collections::{HashMap, HashSet};

use crate::part1::generate_next;

pub fn part2(_input: &str) -> usize {
    let secret_numbers = _input
        .split("\n")
        .filter(|c| !c.is_empty())
        .map(|c| c.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();

    let mut seq_total = HashMap::new();
    for mut secret_num in secret_numbers {
        let mut buyer = vec![secret_num % 10];
        for _ in 0..2000 {
            secret_num = generate_next(secret_num);
            buyer.push(secret_num % 10);
        }
        let mut seen = HashSet::new();
        for i in 0..(buyer.len() - 4) {
            let diff = buyer[i..(i + 5)]
                .windows(2)
                .map(|window| window[1] - window[0])
                .collect::<Vec<i64>>();
            let diff_tup = (diff[0], diff[1], diff[2], diff[3]);
            if seen.contains(&diff_tup) {
                continue;
            }
            seen.insert(diff_tup);
            *seq_total.entry(diff_tup).or_insert(0) += buyer[i + 4];
        }
    }

    return *seq_total.iter().map(|(_, val)| val).max().unwrap() as usize;
}

#[cfg(test)]
mod day22 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = "1\n2\n3\n2024";
        assert_eq!(part2(_input), 23);
    }
}
