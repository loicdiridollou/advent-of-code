//! # Advent of Code - Day 22 - Part One

use cached::proc_macro::cached;

fn mix(new_value: i64, secret_number: i64) -> i64 {
    return new_value ^ secret_number;
}

fn prune(new_value: i64) -> i64 {
    return new_value.rem_euclid(16777216);
}

#[cached]
pub fn generate_next(mut secret_num: i64) -> i64 {
    secret_num = prune(mix(secret_num * 64, secret_num));
    secret_num = prune(mix((secret_num as f64 / 32.0).floor() as i64, secret_num));
    return prune(mix(secret_num * 2048, secret_num));
}

pub fn part1(_input: &str) -> usize {
    let secret_numbers = _input
        .split("\n")
        .filter(|c| !c.is_empty())
        .map(|c| c.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();

    let mut count = 0;
    for mut secret_num in secret_numbers {
        for _ in 0..2000 {
            secret_num = generate_next(secret_num);
        }
        count += secret_num;
    }

    return count as usize;
}

#[cfg(test)]
mod day22 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = "1\n10\n100\n2024";
        assert_eq!(part1(_input), 37327623);
    }
}
