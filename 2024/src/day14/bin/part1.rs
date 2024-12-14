//! # Advent of Code - Day 14 - Part One

use std::collections::HashMap;

use regex::Regex;

#[derive(Debug)]
pub struct Robot {
    pub r: i32,
    pub c: i32,
    dr: i32,
    dc: i32,
}

impl Robot {
    pub fn move_robot(&mut self, n: i32, nr: i32, nc: i32) {
        self.r = (self.r + self.dr * n).rem_euclid(nr);
        self.c = (self.c + self.dc * n).rem_euclid(nc);
    }
}

pub fn parse_robot(robot_str: &str) -> Robot {
    let robot_re = Regex::new(r"p\=(-?\d+),(-?\d+) v\=(-?\d+),(-?\d+)").unwrap();
    let ext = robot_re.captures(&robot_str).unwrap();
    return Robot {
        r: ext[2].parse::<i32>().unwrap(),
        c: ext[1].parse::<i32>().unwrap(),
        dr: ext[4].parse::<i32>().unwrap(),
        dc: ext[3].parse::<i32>().unwrap(),
    };
}

pub fn part1(_input: &str, nr: i32, nc: i32) -> usize {
    let mut robots = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|robot_str| parse_robot(robot_str))
        .collect::<Vec<Robot>>();

    for robot in robots.iter_mut() {
        robot.move_robot(100, nr, nc);
    }

    let mut quadrants = HashMap::new();
    for robot in robots.iter_mut() {
        if robot.r < nr / 2 && robot.c < nc / 2 {
            *quadrants.entry(0).or_insert(0) += 1;
        } else if robot.r < nr / 2 && robot.c > nc / 2 {
            *quadrants.entry(1).or_insert(0) += 1;
        } else if robot.r > nr / 2 && robot.c > nc / 2 {
            *quadrants.entry(2).or_insert(0) += 1;
        } else if robot.r > nr / 2 && robot.c < nc / 2 {
            *quadrants.entry(3).or_insert(0) += 1;
        }
    }

    let mut total = 1;
    for (_, v) in quadrants.iter() {
        total *= v;
    }

    return total;
}

#[cfg(test)]
mod day14 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input, 7, 11), 12);
    }
}
