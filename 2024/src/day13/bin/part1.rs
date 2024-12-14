//! # Advent of Code - Day 13 - Part One

use regex::Regex;

#[derive(Debug)]
pub struct Claw {
    pub xa: f64,
    pub xb: f64,
    pub xprize: f64,
    pub ya: f64,
    pub yb: f64,
    pub yprize: f64,
}

pub fn solve_claw(claw: &Claw) -> [f64; 2] {
    let det = claw.xa * claw.yb - claw.xb * claw.ya;
    if det == 0. {
        return [-1., -1.];
    }
    let a = (claw.xprize * claw.yb - claw.yprize * claw.ya) / det;
    let b = (claw.xa * claw.yprize - claw.xb * claw.xprize) / det;

    return [a, b];
}

fn is_valid_solution(sol: [f64; 2]) -> bool {
    let [a, b] = sol;
    return 0. <= a && a <= 100.0 && a.fract() == 0. && 0. <= b && b <= 100.0 && b.fract() == 0.;
}

pub fn parse_claw(claw: &str, prize_offset: Option<f64>) -> Claw {
    let claw_re = Regex::new(
        r"Button A: X\+(\d*),.Y\+(\d*)\nButton B: X\+(\d*),.Y\+(\d*)\nPrize: X\=(\d+), Y\=(\d+)",
    )
    .unwrap();
    let ext = claw_re.captures(&claw).unwrap();

    return Claw {
        xa: ext[1].parse::<f64>().unwrap(),
        xb: ext[2].parse::<f64>().unwrap(),
        xprize: prize_offset.unwrap_or(0.) + ext[5].parse::<f64>().unwrap(),
        ya: ext[3].parse::<f64>().unwrap(),
        yb: ext[4].parse::<f64>().unwrap(),
        yprize: prize_offset.unwrap_or(0.) + ext[6].parse::<f64>().unwrap(),
    };
}

pub fn part1(_input: &str) -> usize {
    let claws = _input
        .split("\n\n")
        .map(|x| parse_claw(x, None))
        .map(|x| solve_claw(&x))
        .filter(|sol| is_valid_solution(*sol))
        .map(|sol| sol[0] * 3. + sol[1])
        .sum::<f64>();

    return claws as usize;
}

#[cfg(test)]
mod day13 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 480);
    }
}
