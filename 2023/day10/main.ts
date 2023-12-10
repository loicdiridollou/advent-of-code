// Main file to run answers for day 10 of 2023

import fs from "fs";
import { part1 } from "./part1";
import { part2 } from "./part2";

const input = fs.readFileSync("./2023/day10/part1.txt", "utf-8");

console.log("Part 1 result:", part1(input));
console.log("Part 2 result:", part2(input));
