//! # Advent of Code - Day 7 - Part One

use std::i64;

fn try_value(values: Vec<i64>, target: i64) -> bool {
    if values.len() == 1 && target == values[0] {
        return true;
    } else if values.len() == 1 {
        return false;
    } else if values[0] > target {
        return false;
    }

    return try_value(
        vec![vec![values[0] + values[1]], values[2..].to_vec()].concat(),
        target,
    ) || try_value(
        vec![vec![values[0] * values[1]], values[2..].to_vec()].concat(),
        target,
    );
}

pub fn part1(_input: &str) -> usize {
    let new_vec = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|item| {
            item.split(" ")
                .map(|c| c.replace(":", "").parse::<i64>().unwrap())
                .collect::<Vec<i64>>()
        })
        .collect::<Vec<Vec<i64>>>();

    let mut count = 0;
    for item in new_vec.iter() {
        let target = item[0];

        if try_value(item[1..].to_vec(), target) {
            count += target;
        }
    }
    return count as usize;
}

#[cfg(test)]
mod day07 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 3749);
    }
}
