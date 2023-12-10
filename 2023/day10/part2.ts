// Part 2 for day 10 of 2023

export function part2(input: string): number {
  let maze = input
    .split("\n")
    .filter((group) => group != "")
    .map((group) => group.split(""));

  let start: number[] = [];
  for (let row = 0; row < maze.length; row++) {
    for (let col = 0; col < maze[row].length; col++) {
      if (maze[row][col] == "S") {
        start = [row, col];
      }
    }
  }

  let visited = new Set();
  visited.add(`${start[0]}_${start[1]}`);
  let queue = [`${start[0]}_${start[1]}`];
  let maybe_s = new Set(["|", "-", "J", "L", "7", "F"]);

  while (queue.length != 0) {
    let new_pos = queue[0];
    queue = queue.slice(1);
    let r = parseInt(new_pos.split("_")[0]);
    let c = parseInt(new_pos.split("_")[1]);
    let chr = maze[r][c];

    if (
      r > 0 &&
      "S|JL".includes(chr) &&
      "|7F".includes(maze[r - 1][c]) &&
      !visited.has(`${r - 1}_${c}`)
    ) {
      visited.add(`${r - 1}_${c}`);
      queue.push(`${r - 1}_${c}`);
      if (chr == "S") {
        maybe_s = new Set(
          [...maybe_s].filter((i) => new Set(["|", "J", "L"]).has(i)),
        );
      }
    }
    if (
      r < maze.length - 1 &&
      "S|7F".includes(chr) &&
      "|JL".includes(maze[r + 1][c]) &&
      !visited.has(`${r + 1}_${c}`)
    ) {
      visited.add(`${r + 1}_${c}`);
      queue.push(`${r + 1}_${c}`);
      if (chr == "S") {
        maybe_s = new Set(
          [...maybe_s].filter((i) => new Set(["|", "7", "F"]).has(i)),
        );
      }
    }
    if (
      c > 0 &&
      "S-J7".includes(chr) &&
      "-LF".includes(maze[r][c - 1]) &&
      !visited.has(`${r}_${c - 1}`)
    ) {
      visited.add(`${r}_${c - 1}`);
      queue.push(`${r}_${c - 1}`);
      if (chr == "S") {
        maybe_s = new Set(
          [...maybe_s].filter((i) => new Set(["-", "J", "7"]).has(i)),
        );
      }
    }
    if (
      c < maze[0].length - 1 &&
      "S-LF".includes(chr) &&
      "-J7".includes(maze[r][c + 1]) &&
      !visited.has(`${r}_${c + 1}`)
    ) {
      visited.add(`${r}_${c + 1}`);
      queue.push(`${r}_${c + 1}`);
      if (chr == "S") {
        maybe_s = new Set(
          [...maybe_s].filter((i) => new Set(["-", "L", "F"]).has(i)),
        );
      }
    }
  }

  for (let row = 0; row < maze.length; row++) {
    for (let col = 0; col < maze[row].length; col++) {
      if (maze[row][col] == "S") {
        maze[row][col] = maybe_s.values().next().value;
      } else if (!visited.has(`${row}_${col}`)) {
        maze[row][col] = ".";
      }
    }
  }

  let outside = new Set();

  for (let r = 0; r < maze.length; r++) {
    let inside = false;
    let up = null;
    for (let c = 0; c < maze[0].length; c++) {
      let ch = maze[r][c];
      if (ch == "|") {
        inside = !inside;
      } else if ("LF".includes(ch)) {
        up = ch == "L";
      } else if ("7J".includes(ch)) {
        if (ch != (up ? "J" : "7")) {
          inside = !inside;
        }
        up = null;
      }
      if (!inside) {
        outside.add(`${r}_${c}`);
      }
    }
  }

  outside = new Set([...outside, ...visited]);
  return maze.length * maze[0].length - outside.size;
}
