"""Part 1 for day 25 of 2023."""
from pathlib import Path

import networkx as nx


def main() -> None:
    """Solution for part 1."""
    graph = nx.Graph()

    input_data = Path("2023/day25/input.txt").open()
    for line in input_data:
        le, ri = line.split(":")
        for summit in ri.strip().split(" "):
            graph.add_edge(le, summit)
            graph.add_edge(summit, le)

    graph.remove_edges_from(nx.minimum_edge_cut(graph))
    a, b = nx.connected_components(graph)

    print(len(a) * len(b))


if __name__ == "__main__":
    main()
