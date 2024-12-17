//! # Advent of Code - Day 17 - Part Two

use regex::Regex;

use crate::part1::run_program;

fn solver(program: &Vec<i64>, a: i64, b: i64, c: i64, prg_pos: usize) -> usize {
    for i in 0..8 {
        let first_digit_out = run_program(a * 8 + i, b, c, program)[0];
        if first_digit_out == program[prg_pos] {
            if prg_pos == 0 {
                return (a * 8 + i) as usize;
            }
            let e = solver(program, a * 8 + i, b, c, prg_pos - 1);
            if e != 0 {
                return e as usize;
            }
        }
    }
    return 0;
}

pub fn part2(_input: &str) -> usize {
    let re =
        Regex::new(r"Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: (.+)")
            .unwrap();
    let matching = re.captures(_input).unwrap();
    let b = matching[2].to_string().parse::<i64>().unwrap();
    let c = matching[3].to_string().parse::<i64>().unwrap();
    let program = matching[4]
        .split(",")
        .map(|c| c.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();

    return solver(&program, 0, b, c, program.len() - 1);
}

#[cfg(test)]
mod day17 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = "Register A: 2024\nRegister B: 0\nRegister C: 0\n\nProgram: 0,3,5,4,3,0";
        assert_eq!(part2(_input), 117440);
    }
}
