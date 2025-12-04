//! # Advent of Code 202 - Day 01 part 1

pub fn part1(_input: &str) -> usize {
    let v1 = _input.lines().collect::<Vec<&str>>();

    let mut counter: u64 = 0;

    for v2 in v1.iter() {
        let ln = v2.len();
        let mut max_val = 0;
        let mut max_idx = 0;
        for dd in v2[..(ln - 1)].chars().enumerate() {
            if dd.1.to_digit(10).unwrap() > max_val {
                max_val = dd.1.to_digit(10).unwrap();
                max_idx = dd.0;
            }
        }

        let second_chr = v2[(max_idx + 1)..]
            .chars()
            .max()
            .unwrap()
            .to_digit(10)
            .unwrap();
        counter += (max_val * 10 + second_chr) as u64;
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 357);
    }
}
