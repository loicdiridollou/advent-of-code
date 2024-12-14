//! # Advent of Code - Day 13 - Part Two

use crate::part1::{parse_claw, solve_claw};

fn is_valid_solution(sol: [f64; 2]) -> bool {
    let [a, b] = sol;
    return 0. <= a && a.fract() == 0. && 0. <= b && b.fract() == 0.;
}

pub fn part2(_input: &str) -> usize {
    let claws = _input
        .split("\n\n")
        .map(|x| parse_claw(x, Some(10000000000000.)))
        .map(|x| solve_claw(&x))
        .filter(|sol| is_valid_solution(*sol))
        .map(|sol| sol[0] * 3. + sol[1])
        .sum::<f64>();

    return claws as usize;
}

#[cfg(test)]
mod day13 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 875318608908);
    }
}
