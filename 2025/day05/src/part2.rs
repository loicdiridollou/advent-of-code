// Advent of Code 2025 - Day 5 - Part Two
use std::cmp;

pub fn part2(_input: &str) -> usize {
    let v1 = _input.split("\n\n").collect::<Vec<&str>>();
    let mut v2 = v1[0]
        .lines()
        .map(|c| {
            let d = c
                .split('-')
                .map(|b| b.parse::<i64>().unwrap())
                .collect::<Vec<i64>>();
            (d[0], d[1])
        })
        .collect::<Vec<(i64, i64)>>();
    v2.sort_by_key(|&(key, _)| key);
    let mut v3 = vec![v2[0]];

    for item in v2.iter().skip(1) {
        let ln = v3.len();
        if item.0 <= v3[ln - 1].1 {
            v3[ln - 1].1 = cmp::max(v3[ln - 1].1, item.1);
        } else {
            v3.push(*item);
        }
    }
    let counter: i64 = v3.iter().map(|c| c.1 - c.0 + 1).sum();

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 14);
    }
}
