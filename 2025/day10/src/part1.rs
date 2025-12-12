// Advent of Code 2025 - Day 10 part 1

use itertools::Itertools;
use regex::Regex;

fn apply_buttons(mut lights: Vec<bool>, buttons: Vec<usize>) -> Vec<bool> {
    for button in buttons.iter() {
        lights[*button] ^= true;
    }
    lights
}

pub fn part1(_input: &str) -> usize {
    let re = Regex::new(r"\[([#.]+)\] ([()\d, ]+)").unwrap();
    let mut v1: Vec<(Vec<bool>, Vec<Vec<usize>>)> = vec![];
    for vv in _input.lines() {
        let matching = re.captures(vv).unwrap();
        let lights = matching[1].chars().map(|c| c != '.').collect::<Vec<bool>>();
        let buttons = matching[2].to_string();
        let u = buttons
            .split_whitespace()
            .map(|c| {
                c.chars()
                    .filter(|b| *b != '(' && *b != ')')
                    .collect::<String>()
                    .split(',')
                    .map(|c| c.parse::<usize>().unwrap())
                    .collect::<Vec<usize>>()
            })
            .collect::<Vec<Vec<usize>>>();
        v1.push((lights, u));
    }

    let mut v2: Vec<usize> = vec![];
    for (light, buttons) in v1.iter() {
        let mut found = false;
        for n in 1..buttons.len() {
            if found {
                break;
            }
            for bt_list in buttons.iter().combinations(n) {
                if found {
                    break;
                }
                let mut light_clone = vec![false; light.len()];
                for bt in bt_list.iter() {
                    light_clone = apply_buttons(light_clone, bt.to_vec());
                }
                if light_clone == *light {
                    v2.push(n);
                    found = true;
                }
            }
        }
    }

    v2.iter().sum::<usize>()
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 7);
    }
}
