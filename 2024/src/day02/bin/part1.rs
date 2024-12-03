//! # Advent of Code - Day 2 - Part One

pub fn part1(_input: &str) -> usize {
    let v1 = _input
        .lines()
        .map(|x| {
            x.split_whitespace()
                .map(str::to_string)
                .map(|x| x.parse::<i32>().unwrap())
                .collect::<Vec<i32>>()
                .windows(2)
                .map(|slice| (slice[0] - slice[1]))
                .collect::<Vec<i32>>()
        })
        .collect::<Vec<Vec<i32>>>();

    let mut count = 0;

    for el in v1.into_iter() {
        let first_num = el[0];
        if el
            .iter()
            .all(|x| (&first_num * x >= 0) && (x.abs() > 0) && (x.abs() < 4))
        {
            count += 1;
        }
    }

    return count;
}

#[cfg(test)]
mod day02 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 2);
    }
}
