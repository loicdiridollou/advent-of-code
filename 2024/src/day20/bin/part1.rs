//! # Advent of Code - Day 20 - Part One

pub fn part1(_input: &str) -> usize {
    let input = _input
        .split("\n")
        .filter(|c| !c.is_empty())
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut dist = vec![];
    let (mut r, mut c) = (0 as i32, 0 as i32);

    for i in 0..input.len() {
        let mut tmp = vec![];
        for j in 0..input[0].len() {
            if input[i][j] == 'S' {
                r = i as i32;
                c = j as i32;
                tmp.push(0);
            } else {
                tmp.push(-1);
            }
        }
        dist.push(tmp);
    }

    let mut queue = vec![(r, c)];
    while queue.len() > 0 {
        let (cr, cc) = queue.remove(0);
        for (nr, nc) in [(cr + 1, cc), (cr - 1, cc), (cr, cc + 1), (cr, cc - 1)] {
            if nr < 0 || nc < 0 || nr as usize == input.len() || nc as usize == input[0].len() {
                continue;
            }
            if input[nr as usize][nc as usize] == '#' || dist[nr as usize][nc as usize] != -1 {
                continue;
            }
            dist[nr as usize][nc as usize] = dist[cr as usize][cc as usize] + 1;
            queue.push((nr, nc));
        }
    }

    let mut count = 0;

    for r in 0..input.len() {
        for c in 0..input[0].len() {
            if input[r][c] == '#' {
                continue;
            }

            for (nr, nc) in [
                ((r + 2) as i32, c as i32),
                (r as i32 - 1, c as i32 - 1),
                (r as i32 - 1, c as i32 + 1),
                (r as i32 + 1, c as i32 - 1),
                (r as i32 + 1, c as i32 + 1),
                (r as i32 - 2, c as i32),
                (r as i32, c as i32 - 2),
                (r as i32, c as i32 + 2),
            ] {
                if nr < 0 || nc < 0 || nr as usize == input.len() || nc as usize == input[0].len() {
                    continue;
                }
                if input[nr as usize][nc as usize] == '#' {
                    continue;
                }
                if dist[r][c] as i32 - dist[nr as usize][nc as usize] as i32 >= 102 {
                    count += 1;
                }
            }
        }
    }
    return count;
}

#[cfg(test)]
mod day20 {
    use super::*;

    #[test]
    fn test_part1() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part1(_input), 0);
    }
}
