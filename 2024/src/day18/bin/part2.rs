//! # Advent of Code - Day 18 - Part Two

use std::collections::HashMap;

use crate::part1::run_program;

pub fn part2(_input: &str, parameters: (i32, i32, (i32, i32))) -> String {
    let mut mapping = HashMap::new();
    let _input = _input
        .split("\n")
        .filter(|c| !c.is_empty())
        .map(|v| {
            let values = v.split(",").collect::<Vec<&str>>();
            return (
                values[1].parse::<i32>().unwrap(),
                values[0].parse::<i32>().unwrap(),
            );
        })
        .collect::<Vec<(i32, i32)>>();
    for i in 0.._input.len() {
        mapping.insert(_input[i], i);
    }
    let mut l = 0;
    let mut r = _input.len() - 1;
    let mut mid = 0;

    while l < r {
        mid = (l + r) / 2;
        let (aa, _) = run_program(&mapping, parameters, mid + 1);
        let (bb, _) = run_program(&mapping, parameters, mid + 2);

        if aa && !bb {
            break;
        } else if aa && bb {
            l = mid + 1;
        } else {
            r = mid;
        }
    }
    return format!("{},{}", _input[mid + 1].1, _input[mid + 1].0);
}

#[cfg(test)]
mod day18 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = "5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0";
        assert_eq!(part2(_input, (7, 7, (6, 6))), "6,1");
    }
}
