// Advent of Code 2025 - Day 07 - Part Two

pub fn part2(_input: &str) -> usize {
    let v1 = _input
        .lines()
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut curr = vec![0; v1[0].len()];
    for (c, val) in v1[0].iter().enumerate() {
        if *val == 'S' {
            curr[c] = 1;
            break;
        }
    }

    for r in 1..v1.len() {
        let mut new = vec![0; v1[0].len()];
        for c in 0..v1[0].len() {
            if v1[r][c] == '^' {
                new[c - 1] += curr[c];
                new[c + 1] += curr[c];
            } else {
                new[c] += curr[c];
            }
        }
        curr = new;
    }

    curr.iter().sum::<usize>()
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 21);
    }
}
