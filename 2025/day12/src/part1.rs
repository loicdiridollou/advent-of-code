// Advent of Code 2025 - Day 12 part 1

pub fn part1(input: &str) -> usize {
    // 5, 3x3 presents
    // let present_size = 7;
    let presents = parse_presents(input);
    let lines = parse_lines(input);

    lines
        .iter()
        .filter_map(|((x, y), present_counts)| {
            (x * y
                > present_counts
                    .iter()
                    .enumerate()
                    .map(|(index, num_presents)| presents[index].1 as u32 * num_presents)
                    .sum::<u32>())
            .then_some(())
        })
        .count()
}

fn parse_lines(_input: &str) -> Vec<((u32, u32), Vec<u32>)> {
    let v1 = _input.split("\n\n").collect::<Vec<&str>>();
    let dd = v1.iter().last().unwrap();
    let mut res = vec![];
    for uu in dd.lines() {
        let du = uu.split(':').collect::<Vec<&str>>();
        let size = du[0]
            .split('x')
            .map(|c| c.parse::<u32>().unwrap())
            .collect::<Vec<u32>>();
        let config = du[1]
            .split_whitespace()
            .map(|c| c.parse::<u32>().unwrap())
            .collect::<Vec<u32>>();
        res.push(((size[0], size[1]), config));
    }
    res
}

fn parse_presents(_input: &str) -> Vec<(u32, usize)> {
    let v1 = _input.split("\n\n").collect::<Vec<&str>>();
    v1.iter()
        .take(v1.len() - 1)
        .map(|c| {
            let num = c.split(':').collect::<Vec<&str>>()[0]
                .parse::<u32>()
                .unwrap();
            let num_shape = c.split(':').collect::<Vec<&str>>()[1]
                .chars()
                .filter(|v| *v == '#')
                .count();
            (num, num_shape)
        })
        .collect::<Vec<(u32, usize)>>()
}

#[cfg(test)]
mod day01 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 3);
    }
}
