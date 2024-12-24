//! # Advent of Code - Day 24 - Part Two

use std::collections::{HashMap, HashSet};

use regex::Regex;

pub fn part2(_input: &str) -> String {
    let data: Vec<&str> = _input.split("\n\n").collect();
    let mut xs = HashMap::new();
    let mut ys = HashMap::new();
    for l in data[0].lines() {
        let (key, value) = l.split_once(':').unwrap();
        let value = value.trim().parse::<i64>().unwrap();
        if key.contains('x') {
            xs.insert(key, value);
        } else {
            ys.insert(key, value);
        }
    }

    let mut ops: HashMap<&str, (&str, &str, &str)> = HashMap::new();
    let mut rev_ops: HashMap<(&str, &str, &str), &str> = HashMap::new();

    for l in data[1].lines() {
        let parts: Vec<&str> = l.split_whitespace().collect();
        assert!(parts.len() == 5);
        ops.insert(parts[4], (parts[0], parts[1], parts[2]));

        rev_ops.insert((parts[0], parts[1], parts[2]), parts[4]);
        rev_ops.insert((parts[2], parts[1], parts[0]), parts[4]);
    }

    let re = Regex::new(r"z\d+").unwrap();
    let mut top = 0;
    for d in ops.keys() {
        if re.is_match(d) {
            let num = d[1..].parse::<i64>().unwrap();
            top = top.max(num);
        }
    }

    let mut wrong_gates: HashSet<String> = HashSet::new();

    for i in 1..top {
        let x = format!("x{:02}", i);
        let y = format!("y{:02}", i);
        let z = format!("z{:02}", i);

        let res_op = ops.get(&z as &str).unwrap();

        let xor_gate = rev_ops
            .get(&(x.as_str(), "XOR", y.as_str()))
            .unwrap_or(&"emp");
        let and_gate = rev_ops
            .get(&(x.as_str(), "AND", y.as_str()))
            .unwrap_or(&"emp");

        if !res_op.1.contains("XOR") {
            wrong_gates.insert(z.to_string());
        }

        let mut carry = Vec::new();
        for (o0, o1, o2) in ops.values() {
            if o1 == &"XOR" && (o0 == xor_gate || o2 == xor_gate) {
                let mut set = HashSet::from([o0, o1, o2]);
                set.remove(&"XOR");
                set.remove(xor_gate);
                carry.push(set);
            }
        }

        if carry.len() != 1 {
            wrong_gates.insert(xor_gate.to_string());
            wrong_gates.insert(and_gate.to_string());
        } else {
            let carry = carry[0].iter().next().unwrap();
            let xor2_gate = rev_ops.get(&(xor_gate, "XOR", carry)).unwrap();
            if xor2_gate != &z {
                wrong_gates.insert(xor2_gate.to_string());
            }
        }
    }

    let mut wrong_gates: Vec<String> = wrong_gates.into_iter().collect();
    wrong_gates.sort();

    return wrong_gates.join(",");
}

#[cfg(test)]
mod day24 {
    // use super::*;
    //
    // #[test]
    // fn test_part2() {
    //     assert_eq!(part2(""), 0);
    // }
}
