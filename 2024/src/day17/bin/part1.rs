//! # Advent of Code - Day 17 - Part One

use regex::Regex;

fn combo(operand: i64, a: i64, b: i64, c: i64) -> i64 {
    match operand {
        0 => 0,
        1 => 1,
        2 => 2,
        3 => 3,
        4 => a,
        5 => b,
        6 => c,
        _ => panic!("Issue"),
    }
}

pub fn run_program(mut a: i64, mut b: i64, mut c: i64, program: &Vec<i64>) -> Vec<i64> {
    let mut outputs = vec![];
    let mut i = 0;
    while i < program.len() {
        let opcode = program[i];
        let operand = program[i + 1];

        if opcode == 0 {
            a = a >> combo(operand, a, b, c);
        } else if opcode == 1 {
            b = b ^ operand;
        } else if opcode == 2 {
            b = combo(operand, a, b, c) % 8;
        } else if opcode == 3 {
            if a != 0 {
                i = operand as usize;
                continue;
            }
        } else if opcode == 4 {
            b = b ^ c;
        } else if opcode == 5 {
            outputs.push(combo(operand, a, b, c) % 8);
        } else if opcode == 6 {
            b = a >> combo(operand, a, b, c);
        } else if opcode == 7 {
            c = a >> combo(operand, a, b, c);
        }
        i += 2;
    }
    return outputs;
}

pub fn part1(_input: &str) -> String {
    let re =
        Regex::new(r"Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: (.+)")
            .unwrap();
    let matching = re.captures(_input).unwrap();
    let a = matching[1].to_string().parse::<i64>().unwrap();
    let b = matching[2].to_string().parse::<i64>().unwrap();
    let c = matching[3].to_string().parse::<i64>().unwrap();
    let program = matching[4]
        .split(",")
        .map(|c| c.parse::<i64>().unwrap())
        .collect::<Vec<i64>>();

    let outputs = run_program(a, b, c, &program);
    return outputs.iter().map(|&id| id.to_string() + ",").collect();
}

#[cfg(test)]
mod day17 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), "4,6,3,5,6,3,5,2,1,0,");
    }
}
