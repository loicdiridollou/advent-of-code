// Main file to run answers for day 18 of 2023

import fs from "fs";
import { part1 } from "./part1";
import { part2 } from "./part2";

const input = fs.readFileSync("./2023/day18/part1.txt", "utf-8");

let start = Date.now();
console.log("Part 1 result:", part1(input));
console.log(`Part 1 executed in: ${(Date.now() - start) / 1000} s`);
start = Date.now();
console.log("Part 2 result:", part2(input));
console.log(`Part 2 executed in: ${(Date.now() - start) / 1000} s`);
