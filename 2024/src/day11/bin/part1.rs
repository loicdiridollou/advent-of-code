//! # Advent of Code - Day 11 - Part One

use cached::proc_macro::cached;

#[cached]
fn count_stones(stone: i64, steps: i64) -> i64 {
    if steps == 0 {
        return 1;
    } else if stone == 0 {
        return count_stones(1, steps - 1);
    } else if stone.to_string().len() % 2 == 0 {
        let val_str = &stone.to_string();
        let left = val_str[..(val_str.len() / 2)].parse::<i64>().unwrap();
        let right = val_str[(val_str.len() / 2)..].parse::<i64>().unwrap();
        return count_stones(left, steps - 1) + count_stones(right, steps - 1);
    }
    return count_stones(stone * 2024, steps - 1);
}

pub fn part1(_input: &str) -> usize {
    let input = _input
        .split_whitespace()
        .map(|item| item.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();

    let mut count = 0;

    for el in input.iter() {
        count += count_stones(*el, 25);
    }

    return count as usize;
}

#[cfg(test)]
mod day11 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 55312);
    }
}
