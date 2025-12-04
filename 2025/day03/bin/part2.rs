//! # Advent of Code - Day 1 - Part Two

pub fn part2(_input: &str) -> usize {
    let v1 = _input.lines().collect::<Vec<&str>>();

    let mut counter: u64 = 0;
    let num_itm = 12;

    for v2 in v1.iter() {
        let ln = v2.len();
        let mut max_val = 0;
        let mut max_idx = 0;
        for dd in v2[..=(ln - num_itm)].chars().enumerate() {
            if dd.1.to_digit(10).unwrap() > max_val {
                max_val = dd.1.to_digit(10).unwrap();
                max_idx = dd.0;
            }
        }
        let mut cur_counter = max_val as u64;

        for itm in 1..num_itm {
            let curr_idx = max_idx;
            max_val = 0;
            for dd in v2[(max_idx + 1)..=(ln - num_itm + itm)].chars().enumerate() {
                if dd.1.to_digit(10).unwrap() > max_val {
                    max_val = dd.1.to_digit(10).unwrap();
                    max_idx = dd.0 + curr_idx + 1;
                }
            }
            cur_counter = 10 * cur_counter + max_val as u64;
        }
        counter += cur_counter;
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 3121910778619);
    }
}
