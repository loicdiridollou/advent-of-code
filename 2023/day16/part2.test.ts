import fs from "fs";
import { part2 } from "./part2";

const input = fs.readFileSync("./2023/day16/part1.test.txt", "utf-8");
describe("Testing part 2 with a multiple-line string", () => {
  test("Value should be returned for a multiple-line string", () => {
    expect(part2(input)).toBe(51);
  });
});
