// Advent of Code 2025 - Day 06 - Part Two

use std::collections::HashMap;

use pathfinding::prelude::count_paths;

#[derive(Debug, Eq, PartialEq, Hash, Clone, Copy)]
struct Node<'a> {
    label: &'a str,
    fft: bool,
    dac: bool,
}
pub fn part2(_input: &str) -> usize {
    let mut devices: HashMap<&str, Vec<&str>> = HashMap::new();

    _input.lines().for_each(|c| {
        let v = c.split(':').collect::<Vec<&str>>();
        let elem = v[1].split_whitespace().collect::<Vec<&str>>();
        devices.insert(v[0], elem);
    });
    devices.insert("out", vec![]);

    count_paths(
        Node {
            label: "svr",
            fft: false,
            dac: false,
        },
        |&device| {
            devices[device.label].iter().map(move |&next_label| Node {
                label: next_label,
                fft: device.fft || next_label == "fft",
                dac: device.dac || next_label == "dac",
            })
        },
        |&c| {
            c == Node {
                label: "out",
                fft: true,
                dac: true,
            }
        },
    )
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput2.txt");
        assert_eq!(part2(_input), 2);
    }
}
