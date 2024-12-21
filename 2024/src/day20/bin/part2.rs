//! # Advent of Code - Day 20 - Part Two

/// Manhattan distance calculation
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

pub fn part2(_input: &str, max_jump: usize, min_savings: usize) -> usize {
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
    let mut path = vec![(r as usize, c as usize)];
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
            path.push((nr as usize, nc as usize));
            queue.push((nr, nc));
        }
    }

    let mut number_of_cheats = 0;

    for start in 0..path.len() {
        for end in (start + min_savings)..path.len() {
            let manhattan_distance = distance(&path, start, end);
            if manhattan_distance <= max_jump && manhattan_distance + min_savings <= end - start {
                number_of_cheats += 1;
            }
        }
    }

    return number_of_cheats;
}

#[cfg(test)]
mod day20 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input, 3, 12), 29);
    }
}
