// Part 1 for day 08 of 2023

export function createNetwork(data: string[]): { [index: string]: string[] } {
  let res: { [index: string]: string[] } = {};
  data.forEach((group) => {
    let match = group.match(
      /([A-Z0-9]{3})\s=\s\(([A-Z0-9]{3}),\s([A-Z0-9]{3})\)/,
    );
    let dd = match?.slice(1);
    if (dd) {
      res[dd[0]] = [dd[1], dd[2]];
    }
  });
  return res;
}
export function part1(input: string): number {
  let instr = input.split("\n\n")[0];
  let network = createNetwork(input.split("\n\n")[1].split("\n"));
  let start = "AAA";

  let num_iter = 0;

  while (start != "ZZZ") {
    let op = instr[num_iter % instr.length];
    start = network[start][op == "L" ? 0 : 1];
    num_iter++;
  }
  return num_iter;
}
