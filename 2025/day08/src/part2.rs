// Advent of Code 2025 - Day 08 - Part Two

use std::collections::{HashMap, HashSet};

use crate::part1::{convert_to_point, distance, Point};

fn find_connection(
    num: usize,
    v1: &[Point],
    dist_list: Vec<(f64, Point, Point)>,
    mapping: HashMap<Point, usize>,
) -> i64 {
    let mut dist_matrix = vec![vec![0; v1.len()]; v1.len()];
    let (_, mut p1, mut p2): (usize, Point, Point) =
        (0, Point { x: 0, y: 0, z: 0 }, Point { x: 0, y: 0, z: 0 });
    for item in dist_list.iter().take(num) {
        p1 = item.1;
        p2 = item.2;
        dist_matrix[*mapping.get(&p1).unwrap()][*mapping.get(&p2).unwrap()] = 1;
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
    if counter == [v1.len()] {
        p1.x * p2.x
    } else {
        -1
    }
}

pub fn part2(_input: &str) -> usize {
    let v1 = _input.lines().map(convert_to_point).collect::<Vec<Point>>();
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

    let mut l = 2;
    let mut r = dist_list.len();
    // not sure a dichotomy is actually efficient here, or at least in this context
    loop {
        let mid = (l + r) / 2;
        let value = find_connection(mid, &v1, dist_list.clone(), mapping.clone());
        let value_bis = find_connection(mid + 1, &v1, dist_list.clone(), mapping.clone());

        if (value == -1) & (value_bis != -1) {
            return value_bis as usize;
        } else if (value == -1) & (value_bis == -1) {
            l = mid + 1;
        } else {
            r = mid;
        }
    }
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 25272);
    }
}
