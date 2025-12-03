//! # Advent of Code 202 - Day 01 part 1

pub fn part1(_input: &str) -> usize {
    let v1 = _input
        .trim_end()
        .split(',')
        .map(|line| {
            line.split('-')
                .collect::<Vec<&str>>()
                .iter()
                .map(|&b| b.parse::<i64>().unwrap())
                .collect::<Vec<i64>>()
        })
        .collect::<Vec<Vec<i64>>>();

    let mut counter = 0;

    for v2 in v1.iter() {
        for u in v2[0]..=v2[1] {
            let q = u.to_string();
            let ln = q.len() / 2;
            if q.len() % 2 == 0 && q[..ln] == q[ln..] {
                counter += u;
            }
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
        assert_eq!(part1(_input), 1227775554);
    }
}
