// Main file to run answers for day 01 of 2023

import fs from "fs";
import { part1 } from "./part1";
import { part2 } from "./part2";

const input = fs.readFileSync("./2023/day05/part1.test.txt", "utf-8");

console.log("Part 1 result:", part1(input));
console.log("Part 2 result:", part2(input));
