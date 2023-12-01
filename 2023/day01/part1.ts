import fs from "fs";

function is_numeric(str: string) {
  return /^\d+$/.test(str);
}
const words = fs.readFileSync("./2023/day01/part1.txt", "utf-8");
console.log(
  words
    .split("\n")
    .map((group) => group.split("").filter(is_numeric))
    .filter((group) => group.length > 0)
    .map((group) => group[0] + group[group.length - 1])
    .map((val) => parseInt(val))
    .reduce((sum, current) => sum + current, 0),
);
