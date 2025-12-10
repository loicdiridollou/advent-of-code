// Advent of Code 2025 - Day 09 - Part Two

use glam::I64Vec2;
use itertools::Itertools;
use nom::{
    bytes::complete::tag,
    character::complete::{self, line_ending},
    multi::separated_list1,
    sequence::separated_pair,
    IResult, Parser,
};

fn parse(input: &str) -> IResult<&str, Vec<I64Vec2>> {
    separated_list1(
        line_ending,
        separated_pair(complete::i64, tag(","), complete::i64).map(|(x, y)| I64Vec2::new(x, y)),
    )
    .parse(input)
}

pub fn part2(_input: &str) -> usize {
    let (_, red_tiles) = parse(_input).unwrap();
    let lines = red_tiles
        .iter()
        .circular_tuple_windows()
        .collect::<Vec<(&I64Vec2, &I64Vec2)>>();
    let max_box = red_tiles
        .iter()
        .tuple_combinations()
        .map(|(a, b)| {
            let area = (a.x.abs_diff(b.x) + 1) * (a.y.abs_diff(b.y) + 1);
            (a, b, area)
        })
        .sorted_by_key(|v| v.2)
        .rev()
        .find(|(a, b, _area)| {
            lines.iter().all(|(line_start, line_end)| {
                // if line is to left
                let left_of_rect = a.x.max(b.x) <= line_start.x.min(line_end.x);
                let right_of_rect = a.x.min(b.x) >= line_start.x.max(line_end.x);
                let above = a.y.max(b.y) <= line_start.y.min(line_end.y);
                let below = a.y.min(b.y) >= line_start.y.max(line_end.y);
                left_of_rect || right_of_rect || above || below
            })
        });

    max_box.unwrap().2 as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 24);
    }
}
