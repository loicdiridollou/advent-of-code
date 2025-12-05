// Advent of Code 2025 - Day 04 part 1

pub fn part1(_input: &str) -> usize {
    let v1 = _input
        .lines()
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut counter = 0;
    for i in 0..v1.len() {
        for j in 0..v1[0].len() {
            if v1[i][j] != '@' {
                continue;
            }

            let mut n = true;
            let mut nw = true;
            let mut ne = true;
            let mut w = true;
            let mut sw = true;
            let mut s = true;
            let mut se = true;
            let mut e = true;

            if i > 0 && v1[i - 1][j] == '@' {
                n = false;
            }
            if i > 0 && j > 0 && v1[i - 1][j - 1] == '@' {
                nw = false;
            }
            if i > 0 && j < v1[0].len() - 1 && v1[i - 1][j + 1] == '@' {
                ne = false;
            }
            if j > 0 && v1[i][j - 1] == '@' {
                w = false;
            }
            if j < v1[0].len() - 1 && v1[i][j + 1] == '@' {
                e = false;
            }
            if i < v1[0].len() - 1 && v1[i + 1][j] == '@' {
                s = false;
            }
            if i < v1[0].len() - 1 && j > 0 && v1[i + 1][j - 1] == '@' {
                sw = false;
            }
            if i < v1[0].len() - 1 && j < v1[0].len() - 1 && v1[i + 1][j + 1] == '@' {
                se = false;
            }

            let v3: i32 = vec![n, nw, ne, w, e, sw, s, se]
                .into_iter()
                .map(|c| if c { 1 } else { 0 })
                .sum();
            if v3 > 4 {
                counter += 1;
            }
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
        assert_eq!(part1(_input), 43);
    }
}
