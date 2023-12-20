// Part 2 for day 20 of 2023

var gcd = function (a: number, b: number) {
  if (!b) {
    return a;
  }

  return gcd(b, a % b);
};
import { Module } from "./part1";

export function part2(input: string): number {
  let modules: { [index: string]: Module } = {};
  let broadcastTargets: string[] = [];

  for (let line of input.split("\n")) {
    if (line == "") {
      continue;
    }

    let [left, right] = line.split(" -> ");
    let outputs = right.split(", ");
    if (left == "broadcaster") {
      broadcastTargets = outputs;
    } else {
      let type = left[0];
      let name = left.slice(1);
      modules[name] = new Module(name, type, outputs);
    }
  }

  for (let key in modules) {
    for (let output of modules[key].outputs) {
      if (
        Object.keys(modules).includes(output) &&
        modules[output].type == "&"
      ) {
        modules[output].memory[key] = "lo";
      }
    }
  }

  let feeds = [];
  for (let key in modules) {
    if (modules[key].outputs.includes("rx")) {
      feeds.push(key);
    }
  }

  let cycleLengths: { [index: string]: number } = {};
  let seen: { [index: string]: number } = {};
  for (let key in modules) {
    if (modules[key].outputs.includes(feeds[0])) {
      seen[key] = 0;
    }
  }
  let presses = 0;

  while (true) {
    presses += 1;
    let queue = [];
    for (let tgt of broadcastTargets) {
      queue.push(["broadcaster", tgt, "lo"]);
    }

    while (queue.length > 0) {
      let [origin, target, pulse] = queue[0];
      queue = queue.slice(1);

      if (!Object.keys(modules).includes(target)) {
        continue;
      }

      let module = modules[target];

      if (module.name == feeds[0] && pulse == "hi") {
        seen[origin] += 1;

        if (!Object.keys(cycleLengths).includes(origin)) {
          cycleLengths[origin] = presses;
        }

        if (Object.values(seen).every((c) => c != 0)) {
          let x = 1;
          for (let cycle of Object.values(cycleLengths)) {
            x = (x * cycle) / gcd(cycle, x);
          }
          return x;
        }
      }

      if (module.type == "%") {
        if (pulse == "lo") {
          module.memory = module.memory == "on" ? "off" : "on";
          let outgoing = module.memory == "on" ? "hi" : "lo";
          for (let output of module.outputs) {
            queue.push([module.name, output, outgoing]);
          }
        }
      } else {
        module.memory[origin] = pulse;
        let outgoing = Object.values(module.memory).every((c) => c == "hi")
          ? "lo"
          : "hi";
        for (let output of module.outputs) {
          queue.push([module.name, output, outgoing]);
        }
      }
    }
  }

  return -1;
}
