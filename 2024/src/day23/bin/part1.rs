//! # Advent of Code - Day 23 - Part One

use std::collections::HashMap;
use std::collections::HashSet;

pub fn part1(_input: &str) -> usize {
    let mut graph = HashMap::new();
    _input.split("\n").filter(|c| !c.is_empty()).for_each(|x| {
        let val = x.split("-").collect::<Vec<&str>>();
        graph.entry(val[0]).or_insert(HashSet::new()).insert(val[1]);
        graph.entry(val[1]).or_insert(HashSet::new()).insert(val[0]);
    });

    let mut connected_sets = HashSet::new();

    for (key, list_val) in graph.iter() {
        let a = key;
        for el in list_val.iter() {
            let b = el;
            for c in graph.get(el).unwrap() {
                if !graph.get(c).unwrap().contains(a) {
                    continue;
                }
                if a == b || b == c || a == c {
                    continue;
                }
                if a[..1] == *"t" || b[..1] == *"t" || c[..1] == *"t" {
                    let mut v = [a, b, c];
                    v.sort();
                    connected_sets.insert(v);
                }
            }
        }
    }
    return connected_sets.len();
}

#[cfg(test)]
mod day23 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 7);
    }
}
