//! # Advent of Code - Day 24 - Part One

use std::collections::HashMap;

pub fn apply_command(
    item: &str,
    commands: &HashMap<&str, (&str, &str, &str)>,
    known_values: &HashMap<&str, i32>,
) -> i32 {
    let (op, x, y) = commands[&item];
    let new_x: i32;
    let new_y: i32;
    if known_values.get(x).is_none() {
        new_x = apply_command(x, &commands, &known_values);
    } else {
        new_x = *known_values.get(x).unwrap();
    }
    if known_values.get(y).is_none() {
        new_y = apply_command(y, &commands, &known_values);
    } else {
        new_y = *known_values.get(y).unwrap();
    }
    return match op {
        "AND" => new_x & new_y,
        "OR" => new_x | new_y,
        "XOR" => new_x ^ new_y,
        _ => panic!("Wrong value of {op}"),
    };
}

pub fn part1(_input: &str) -> usize {
    let values = _input.split("\n\n").collect::<Vec<&str>>();
    let mut known_values = HashMap::new();
    values[0]
        .split("\n")
        .filter(|c| !c.is_empty())
        .for_each(|x| {
            let val = x.split(": ").collect::<Vec<&str>>();
            let (a, b) = (val[0], val[1]);
            known_values.insert(a, b.parse::<i32>().unwrap());
        });
    let mut commands = HashMap::new();
    values[1]
        .split("\n")
        .filter(|c| !c.is_empty())
        .for_each(|x| {
            let val = x.split(" ").collect::<Vec<&str>>();
            let (x, op, y, z) = (val[0], val[1], val[2], val[4]);
            commands.insert(z, (op, x, y));
        });

    let mut results = String::new();
    let mut idx = 0;
    loop {
        let curr = format!("z{:0>2}", idx);
        if commands.get(&curr.as_str()).is_none() {
            break;
        }
        let value = apply_command(curr.as_str(), &commands, &known_values);
        results = value.to_string() + &results;
        idx += 1;
    }

    return i64::from_str_radix(&results, 2).expect("Not a binary number!") as usize;
}

#[cfg(test)]
mod day24 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 9);
    }
}
