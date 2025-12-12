// Advent of Code 2025 - Day 12

use std::{fs::read_to_string, time::Instant};

mod part1;

fn main() {
    let _input = read_to_string("./input.txt").unwrap();

    // start timer
    let start = Instant::now();

    // count and print
    println!("Part 1: {}", part1::part1(&_input));

    // print time taken by part1
    println!("Time taken by Part 1: {:?}", start.elapsed());
}
