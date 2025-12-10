// Advent of Code 2025 - Day 09 part 1

#[derive(Debug, Hash, Eq, PartialEq)]
pub struct Point {
    pub r: i64,
    pub c: i64,
}

pub fn convert_to_point(val: &str) -> Point {
    let dd = val.split(',').collect::<Vec<&str>>();
    Point {
        r: dd[0].parse::<i64>().unwrap(),
        c: dd[1].parse::<i64>().unwrap(),
    }
}

fn compute_area(p1: &Point, p2: &Point) -> i64 {
    ((p1.r - p2.r).abs() + 1) * ((p1.c - p2.c).abs() + 1)
}

pub fn part1(_input: &str) -> usize {
    let points = _input.lines().map(convert_to_point).collect::<Vec<Point>>();
    let mut max_area = 0;

    for (i, p1) in points.iter().enumerate() {
        for p2 in points.iter().skip(i + 1) {
            let area = compute_area(p1, p2);
            max_area = std::cmp::max(area, max_area);
        }
    }
    max_area as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 50);
    }
}
