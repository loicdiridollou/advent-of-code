// Advent of Code 2025 - Day 06 part 1

pub fn part1(_input: &str) -> usize {
    let mut v3: Vec<Vec<i64>> = vec![];
    let mut v4: Vec<&str> = vec![];

    let ln = _input.lines().count();

    for (i, v) in _input.lines().enumerate() {
        let u = v.split_whitespace();
        if i == ln - 1 {
            v4 = u.collect();
        } else {
            v3.push(u.map(|c| c.parse::<i64>().unwrap()).collect::<Vec<i64>>());
        }
    }

    let mut counter = 0;
    for (i, elem) in v4.iter().enumerate() {
        let vv = v3.iter().map(|c| c[i]).collect::<Vec<i64>>();
        if *elem == "+" {
            counter += vv.iter().sum::<i64>();
        } else if *elem == "*" {
            counter += vv.iter().product::<i64>();
        }
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 4277556);
    }
}
