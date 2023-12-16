import { part1 } from "./part1";
import fs from "fs";

const input = fs.readFileSync("./2023/day16/part1.test.txt", "utf-8");

describe("Testing part 1 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(part1(input)).toBe(46);
  });
});
