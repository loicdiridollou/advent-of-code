// Advent of Code 2025 - Day 010 - Part Two

use std::collections::HashSet;

use nom::{
    branch::alt,
    character::complete::{self, line_ending, space1},
    multi::{fold_many1, separated_list1},
    sequence::delimited,
    IResult, Parser,
};
use rayon::prelude::*;

pub fn part2(input: &str) -> usize {
    let (_, machines) = machines(input).unwrap();
    machines
        .par_iter()
        .enumerate()
        .map(|(_id, machine)| {
            let mut set = HashSet::<Vec<usize>>::new();
            set.insert(machine.state.clone());
            let mut i = 0;
            loop {
                set = set
                    .into_iter()
                    .flat_map(|state| {
                        machine
                            .buttons
                            .iter()
                            .map(move |button| push_button(state.clone(), button))
                    })
                    .filter(|state| state.iter().zip(machine.jolt.iter()).all(|(a, b)| a <= b))
                    .collect();
                i += 1;
                if set.contains(&machine.jolt) {
                    break;
                }
            }

            i
        })
        .sum::<usize>()
}

#[derive(Debug)]
struct Machine {
    state: Vec<usize>,
    buttons: Vec<Vec<usize>>,
    jolt: Vec<usize>,
}

// use memoize::memoize;

// #[memoize]
fn push_button(mut state: Vec<usize>, button: &[usize]) -> Vec<usize> {
    for bit in button {
        state[*bit] += 1;
    }
    state
}

// use bitvec::prelude::*;

// let mut bv: BitVec = BitVec::new();
// bv.push(false);
// bv.push(true);
fn machines(input: &str) -> IResult<&str, Vec<Machine>> {
    separated_list1(line_ending, machine).parse(input)
}
fn machine(input: &str) -> IResult<&str, Machine> {
    let (input, _goal_set) = goal(input)?;
    let (input, _) = space1(input)?;
    let (input, buttons) = separated_list1(space1, button).parse(input)?;
    let (input, _) = space1(input)?;
    let (input, jolt) = joltage(input)?;
    Ok((
        input,
        Machine {
            state: vec![0; jolt.len()],
            buttons,
            jolt,
        },
    ))
}
// [.##.]
fn goal(input: &str) -> IResult<&str, Vec<usize>> {
    delimited(
        complete::char('['),
        fold_many1(
            alt((complete::char('.'), complete::char('#'))),
            Vec::new,
            |mut acc: Vec<_>, item| {
                acc.push(match item {
                    '.' => 0,
                    '#' => 1,
                    _ => {
                        panic!("invalid!");
                    }
                });
                acc
            },
        ),
        complete::char(']'),
    )
    .parse(input)
}
fn button(input: &str) -> IResult<&str, Vec<usize>> {
    delimited(
        complete::char('('),
        separated_list1(complete::char(','), complete::usize),
        complete::char(')'),
    )
    .parse(input)
}
fn joltage(input: &str) -> IResult<&str, Vec<usize>> {
    delimited(
        complete::char('{'),
        separated_list1(complete::char(','), complete::usize),
        complete::char('}'),
    )
    .parse(input)
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 3263827);
    }
}
