//! # Advent of Code - Day 23 - Part Two
//!
use itertools::Itertools;
use petgraph::prelude::UnGraphMap;

pub fn part2(_input: &str, val: usize) -> String {
    let edges = _input
        .split("\n")
        .filter(|c| !c.is_empty())
        .map(|x| {
            let val = x.split("-").collect::<Vec<&str>>();
            return (val[0], val[1]);
        })
        .collect::<Vec<(&str, &str)>>();

    let g = &UnGraphMap::<&str, ()>::from_edges(&edges);

    let output = g
        .nodes()
        .flat_map(|node| {
            g.neighbors(node)
                .combinations(val)
                .filter_map(move |neighbor_subset| {
                    if neighbor_subset
                        .iter()
                        .tuple_combinations()
                        .all(move |(a, b)| g.contains_edge(a, b))
                    {
                        let mut nodes = vec![node]
                            .into_iter()
                            .chain(neighbor_subset.into_iter())
                            .collect::<Vec<_>>();
                        nodes.sort();
                        Some(nodes)
                    } else {
                        None
                    }
                })
        })
        .unique()
        .collect::<Vec<_>>();

    return output[0].join(",").to_string();
}

#[cfg(test)]
mod day23 {
    use super::*;

    #[test]
    fn test_part2() {
        let _input = include_str!("../testinput.txt");
        assert_eq!(part2(_input, 3), "co,de,ka,ta");
    }
}
