//! # Advent of Code - Day 20 - Part Two

fn distance(path: &Vec<(usize, usize)>, start: usize, end: usize) -> usize {
    return path
        .get(start)
        .unwrap()
        .0
        .abs_diff(path.get(end).unwrap().0)
        + path
            .get(start)
            .unwrap()
            .1
            .abs_diff(path.get(end).unwrap().1);
}

fn find_char(map: &Vec<Vec<char>>, value: char) -> (usize, usize) {
    map.iter()
        .enumerate()
        .flat_map(|(y, line)| line.iter().enumerate().map(move |(x, c)| (x, y, c)))
        .find(|c| *c.2 == value)
        .map(|c| (c.0, c.1))
        .unwrap()
}

pub fn read_path(input: &str) -> Vec<(usize, usize)> {
    let map = input
        .lines()
        .map(|l| l.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    let start = find_char(&map, 'S');
    let end = find_char(&map, 'E');
    let mut path = Vec::<(usize, usize)>::new();
    path.push(start);
    let mut current = start;
    'parse_map: loop {
        for next in [
            (current.0 - 1, current.1),
            (current.0 + 1, current.1),
            (current.0, current.1 - 1),
            (current.0, current.1 + 1),
        ] {
            if path.iter().rev().nth(1) != Some(&next) && map[next.1][next.0] != '#' {
                path.push(next);
                if next == end {
                    break 'parse_map;
                }
                current = next;
                break;
            }
        }
    }
    return path;
}

pub fn part2(_input: &str) -> i32 {
    let path = read_path(_input);
    let mut number_of_cheats = 0;

    for start in 0..path.len() {
        for end in (start + 100)..path.len() {
            let manhattan_distance = distance(&path, start, end);
            if manhattan_distance <= 20 && manhattan_distance <= end - start - 100 {
                number_of_cheats += 1;
            }
        }
    }

    return number_of_cheats;
}

#[cfg(test)]
mod day20 {
    use super::*;
    //
    // #[test]
    // fn test_part2() {
    //     assert_eq!(part2(""), 0);
    // }
}
