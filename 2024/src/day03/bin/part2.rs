//! # Advent of Code - Day 3 - Part Two

use regex::Regex;

pub fn part2(_input: &str) -> usize {
    let mul_re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();
    let flp_re = Regex::new(r"(mul\((\d+),(\d+)\)|do\(\)|don't\(\))").unwrap();

    let mut mul_active = true;
    let mut count = 0;

    for val in flp_re.captures_iter(_input).map(|c| c.get(0)) {
        if val.unwrap().as_str() == "do()" {
            mul_active = true;
        } else if val.unwrap().as_str() == "don't()" {
            mul_active = false;
        } else if mul_active {
            let extract = mul_re.captures(&val.unwrap().as_str()).unwrap();
            count += extract[1].parse::<i32>().unwrap() * extract[2].parse::<i32>().unwrap();
        }
    }
    return count as usize;
}

#[cfg(test)]
mod day03 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";
        assert_eq!(part2(_input), 48);
    }
}
