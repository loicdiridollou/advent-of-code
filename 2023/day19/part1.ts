// Part 1 for day 19 of 2023

function parseParts(input: string): { [index: string]: number } {
  const regexpSize = /{x=(-?[0-9]*),m=(-?\d*),a=(-?\d*),s=(-?\d*)}/;
  const match = input.match(regexpSize);
  return {
    x: parseInt(match![1]),
    m: parseInt(match![2]),
    a: parseInt(match![3]),
    s: parseInt(match![4]),
  };
}

export function parseInstruction(input: string): [string, any[]] {
  let cnt = input.match(/(.*){(.*)}/);

  let split = cnt![2].split(",");
  let term = split[split.length - 1];
  let instr = [];

  for (let idd of split.slice(0, split.length - 1)) {
    let match = idd.match(/(.*)([<>])(.*):(.*)/);
    if (match) {
      instr.push(match.slice(1));
    }
  }
  instr.push(term);

  return [cnt![1], instr];
}

export function part1(input: string): number {
  let [instrRaw, partsRaw] = input
    .split("\n\n")
    .filter((c) => c != "")
    .map((c) => c.split("\n").filter((c) => c != ""));

  let instrMap: { [index: string]: any[] } = {};
  for (let instr of instrRaw) {
    let [name, ins] = parseInstruction(instr);
    instrMap[name] = ins;
  }

  let parts: { [index: string]: number }[] = partsRaw.map((c) => parseParts(c));
  let total = 0;

  for (let part of parts) {
    let pos = "in";
    while (true) {
      if (pos == "A") {
        total += part["x"] + part["m"] + part["a"] + part["s"];
        break;
      } else if (pos == "R") {
        break;
      }
      let currInstr = instrMap[pos];
      for (let el of currInstr) {
        if (typeof el === "string") {
          pos = el;
          break;
        }
        if (el[1] == ">") {
          if (part[el[0]] > parseInt(el[2])) {
            pos = el[3];
            break;
          }
        } else {
          if (part[el[0]] < parseInt(el[2])) {
            pos = el[3];
            break;
          }
        }
      }
    }
  }

  return total;
}
