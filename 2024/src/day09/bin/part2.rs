//! # Advent of Code - Day 9 - Part Two

fn clean_up_values(mut values: Vec<Vec<i64>>) -> Vec<Vec<i64>> {
    let mut i = 0;

    while i <= values.len() - 2 {
        if values[i][0] == values[i + 1][0] && values[i][0] == -1 {
            values[i][1] = values[i][1] + values[i + 1][1];
            values.remove(i + 1);
        }
        i += 1;
    }
    return values;
}

fn move_block(mut values: Vec<Vec<i64>>, id: i64) -> Vec<Vec<i64>> {
    let mut l = 0;
    let mut r = values.len() - 1;

    while l < r {
        if values[l][0] != -1 {
            l += 1;
        } else if values[r][0] != id || values[r][0] == -1 {
            r -= 1;
        } else if values[r][0] == id && values[l][1] >= values[r][1] {
            let new_bloc_remain = vec![values[r][0], values[r][1]];
            let remaining = values[l][1] - values[r][1];
            values[l] = new_bloc_remain;
            values[r] = vec![-1, values[r][1]];
            if remaining > 0 {
                values.insert(l + 1, vec![-1, remaining]);
            }
            break;
        } else if values[l][0] == -1 && values[l][1] < values[r][1] {
            l += 1;
        }
    }
    return clean_up_values(values);
}
pub fn part2(_input: &str) -> usize {
    let mut values: Vec<Vec<i64>> = vec![];
    let mut id = 0;

    for (idx, el) in _input.chars().enumerate() {
        if el == '\n' {
            continue;
        } else if idx % 2 == 0 {
            values.push(vec![id, el.to_string().parse::<i64>().unwrap()]);
            id += 1;
        } else {
            values.push(vec![-1, el.to_string().parse::<i64>().unwrap()]);
        }
    }

    for tmp_id in (0..id).rev() {
        values = move_block(values, tmp_id);
    }
    let mut count = 0;
    let mut pos = 0;
    for el in values.iter() {
        if el[0] == -1 {
            pos += el[1];
            continue;
        }
        for _ in 1..=el[1] {
            count += el[0] * pos;
            pos += 1
        }
    }
    return count as usize;
}

#[cfg(test)]
mod day09 {
    use super::*;

    #[test]
    fn test_part2() {
        assert_eq!(part2(""), 0);
    }
}
