// Advent of Code 2025 - Day 08 part 1

use std::collections::{HashMap, HashSet};

#[derive(Debug, Clone, Copy, Hash, Eq, PartialEq)]
pub struct Point {
    pub x: i64,
    pub y: i64,
    pub z: i64,
}

pub fn convert_to_point(val: &str) -> Point {
    let dd = val.split(',').collect::<Vec<&str>>();
    Point {
        x: dd[0].parse::<i64>().unwrap(),
        y: dd[1].parse::<i64>().unwrap(),
        z: dd[2].parse::<i64>().unwrap(),
    }
}

pub fn distance(p1: &Point, p2: &Point) -> f64 {
    (((p1.x - p2.x).pow(2) + (p1.y - p2.y).pow(2) + (p1.z - p2.z).pow(2)) as f64).sqrt()
}

pub fn part1(_input: &str, num_boxes: usize) -> usize {
    let v1 = _input.lines().map(convert_to_point).collect::<Vec<Point>>();
    let mut dist_matrix = vec![vec![0; v1.len()]; v1.len()];
    let mut dist_list: Vec<(f64, Point, Point)> = vec![];
    let mut mapping: HashMap<Point, usize> = HashMap::new();

    for (i, p) in v1.iter().enumerate() {
        mapping.insert(*p, i);
    }

    for i in 0..v1.len() - 1 {
        for j in i + 1..v1.len() {
            dist_list.push((distance(&v1[i], &v1[j]), v1[i], v1[j]));
        }
    }

    dist_list.sort_unstable_by(|a, b| a.0.partial_cmp(&b.0).unwrap());

    for item in dist_list.iter().take(num_boxes) {
        let (_, p1, p2) = item;
        dist_matrix[*mapping.get(p1).unwrap()][*mapping.get(p2).unwrap()] = 1;
    }

    let mut connections: HashMap<usize, Vec<usize>> = HashMap::new();

    for (r, row) in dist_matrix.iter().enumerate() {
        for (c, value) in row.iter().enumerate().skip(r + 1) {
            if *value == 1 {
                let vv = connections.entry(r).or_default();
                vv.push(c);
                let vv = connections.entry(c).or_default();
                vv.push(r);
            }
        }
    }

    let mut visited: HashSet<usize> = HashSet::new();
    let mut counter: Vec<usize> = vec![];

    for (key, _) in connections.iter() {
        let mut v3: Vec<usize> = vec![];
        if visited.contains(key) {
            continue;
        }
        let mut queue = vec![key];
        while !queue.is_empty() {
            let curr = queue.remove(0);
            if visited.contains(curr) {
                continue;
            } else {
                v3.push(*curr);
                visited.insert(*curr);
                for u in connections.get(curr).unwrap() {
                    queue.push(u);
                }
            }
        }
        counter.push(v3.len());
    }
    counter.sort();
    counter.iter().skip(counter.len() - 3).product::<usize>()
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input, 10), 40);
    }
}
