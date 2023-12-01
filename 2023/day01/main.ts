// Main file to run answers for day 01 of 2023

import fs from "fs";
import { part1 } from "./part1";
import { part2 } from "./part2";

const words = fs.readFileSync("./2023/day01/part1.txt", "utf-8");

console.log("Part 1 result:", part1(words));
console.log("Part 1 result:", part2(words));
