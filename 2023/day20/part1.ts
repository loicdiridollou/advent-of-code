// Part 1 for day 20 of 2023

export class Module {
  name: string;
  type: string;
  memory: any;
  outputs: string[];

  constructor(name: string, type: string, outputs: string[]) {
    this.name = name;
    this.type = type;
    this.outputs = outputs;

    if (this.type == "%") {
      this.memory = "off";
    } else {
      this.memory = {};
    }
  }
}

export function part1(input: string): number {
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

  let [lo, hi] = [0, 0];

  for (let i = 0; i < 1000; i++) {
    lo += 1;
    let queue = [];
    for (let tgt of broadcastTargets) {
      queue.push(["broadcaster", tgt, "lo"]);
    }

    while (queue.length > 0) {
      let [origin, target, pulse] = queue[0];
      queue = queue.slice(1);

      if (pulse == "lo") {
        lo += 1;
      } else {
        hi += 1;
      }

      if (!Object.keys(modules).includes(target)) {
        continue;
      }

      let module = modules[target];

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

  return lo * hi;
}
