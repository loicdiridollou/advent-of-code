//! # Advent of Code - Day 4

use std::time::Instant;

mod part1;
mod part2;

fn main() {
    let _input = include_str!("../input.txt");

    // start timer
    let start = Instant::now();

    // count and print
    println!("Part 1: {}", part1::part1(&_input));

    // print time taken by part1
    println!("Time taken by Part 1: {:?}", start.elapsed());

    // reset timer
    let start = Instant::now();

    println!("Part 2: {}", part2::part2(&_input));

    // print time taken by part2
    println!("Time taken by Part 2: {:?}", start.elapsed());
}
