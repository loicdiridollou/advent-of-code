"""Solution for day24 of 2023."""

from pathlib import Path

import sympy

hailstones = [
    tuple(map(int, val.replace("@", ",").split(",")))
    for val in Path("2023/day24/input.txt").open()
]

# prepare intersection values
xi, yi, zi, vxi, vyi, vzi = sympy.symbols("xi, yi, zi, vxi, vyi, vzi")

# get equations list to solve
equations = []

for i, (x, y, z, vx, vy, vz) in enumerate(hailstones):
    equations.append((xi - x) * (vy - vyi) - (yi - y) * (vx - vxi))
    equations.append((yi - y) * (vz - vzi) - (zi - z) * (vy - vyi))
    if i < 2:
        continue
    solves = [
        soln
        for soln in sympy.solve(equations)
        if all(x % 1 == 0 for x in soln.values())
    ]
    if len(solves) == 1:
        break

solution = solves[0]

print(solution[xi] + solution[yi] + solution[zi])
