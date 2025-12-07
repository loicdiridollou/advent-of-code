// Advent of Code 2025 - Day 06 - Part Two

pub fn part2(_input: &str) -> usize {
    let v1 = _input
        .lines()
        .map(|c| c.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    let ln = v1.len();
    let mut counter: u64 = 0;

    let mut v4: Vec<u64> = vec![];
    for cc in 0..v1[0].len() {
        let c = v1[0].len() - cc - 1;
        let opr = v1[ln - 1][c];
        let mut v3: Vec<u64> = vec![];
        for item in v1.iter().take(v1.len() - 1) {
            let chr = item[c];
            if chr != ' ' {
                v3.push(chr.to_digit(10).unwrap() as u64);
            }
        }
        let mut cnt = 0;
        if v3.is_empty() {
            continue;
        }
        for (i, vv) in v3.iter().rev().enumerate() {
            cnt += vv * 10u64.pow(i as u32);
        }

        v4.push(cnt);
        if opr == '+' {
            counter += v4.iter().sum::<u64>();
            v4 = vec![];
        } else if opr == '*' {
            counter += v4.iter().product::<u64>();
            v4 = vec![];
        }
    }

    counter as usize
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input), 3263827);
    }
}
