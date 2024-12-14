//! # Advent of Code - Day 14 - Part Two

use std::collections::HashMap;

use crate::part1::{parse_robot, Robot};

pub fn part2(_input: &str, nr: i32, nc: i32) -> usize {
    let mut robots = _input
        .split("\n")
        .filter(|x| !x.is_empty())
        .map(|robot_str| parse_robot(robot_str))
        .collect::<Vec<Robot>>();

    let mut best_iter: i64 = -1;
    let mut min_sf = nr * nc * nr * nc;
    for i in 0..(nr * nc) {
        for robot in robots.iter_mut() {
            robot.move_robot(1, nr, nc);
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
        if total < min_sf {
            min_sf = total;
            best_iter = i as i64;
        }
    }

    return best_iter as usize + 1;
}

#[cfg(test)]
mod day14 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input, 7, 11), 7);
    }
}
