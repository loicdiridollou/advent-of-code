//! # Advent of Code - Day 2 - Part Two

fn check_valid(el: Vec<i32>) -> bool {
    let new_el: Vec<i32> = el.windows(2).map(|slice| (slice[0] - slice[1])).collect();
    let first_num = new_el[0];
    return new_el
        .iter()
        .all(|x| (&first_num * x >= 0) && (x.abs() > 0) && (x.abs() < 4));
}

pub fn part2(_input: &str) -> usize {
    let v1 = _input
        .lines()
        .map(|x| {
            x.split_whitespace()
                .map(str::to_string)
                .map(|x| x.parse::<i32>().unwrap())
                .collect::<Vec<i32>>()
        })
        .collect::<Vec<Vec<i32>>>();

    let mut count = 0;

    for el in v1.into_iter() {
        let n = el.len();
        if check_valid(el.clone()) {
            count += 1;
        } else {
            for i in 0..n {
                let mut new_vec = el.clone();
                new_vec.remove(i);
                if check_valid(new_vec.to_vec()) {
                    count += 1;
                    break;
                }
            }
        }
    }

    return count;
}

#[cfg(test)]
mod day02 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 4);
    }
}
