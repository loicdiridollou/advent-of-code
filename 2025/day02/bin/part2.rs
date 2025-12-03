//! # Advent of Code - Day 1 - Part Two

pub fn part2(_input: &str) -> usize {
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
            for num in 1..=ln {
                let prefix = &q.as_bytes()[..num];
                if q.as_bytes().chunks(num).all(|a| a == prefix) {
                    counter += u;
                    break;
                }
            }
        }
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 4174379265);
    }
}
