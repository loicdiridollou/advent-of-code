// Part 1 for day 17 of 2023

export function heappush(heap: number[][], newKey: number[]) {
  heap.push(newKey);
  let curr = heap.length - 1;

  while (curr > 0) {
    let parent = Math.floor((curr - 1) / 2);
    if (heap[curr][0] < heap[parent][0]) {
      [heap[curr], heap[parent]] = [heap[parent], heap[curr]];
      curr = parent;
    } else {
      break;
    }
  }
}

export function heappop(heap: number[][]): number[] {
  const n = heap.length;
  [heap[0], heap[n - 1]] = [heap[n - 1], heap[0]];
  const removedKey = heap.pop() as number[];

  let curr = 0;

  while (2 * curr + 1 < heap.length) {
    const leftIndex = 2 * curr + 1;
    const rightIndex = 2 * curr + 2;
    const minChildIndex =
      rightIndex < heap.length && heap[rightIndex][0] < heap[leftIndex][0]
        ? rightIndex
        : leftIndex;
    if (heap[minChildIndex][0] < heap[curr][0]) {
      [heap[minChildIndex], heap[curr]] = [heap[curr], heap[minChildIndex]];
      curr = minChildIndex;
    } else {
      break;
    }
  }

  // finally return the removed key
  return removedKey;
}

export function part1(input: string): number {
  let grid = input
    .split("\n")
    .filter((c) => c != "")
    .map((c) => c.split("").map((c) => parseInt(c)));

  let hq = [[0, 0, 0, 0, 0, 0]];
  let seen = new Set<string>();
  let total = -1;

  while (hq.length > 0) {
    let curr = heappop(hq);
    let [hl, r, c, dr, dc, n] = curr;

    if (r == grid.length - 1 && c == grid[0].length - 1) {
      total = hl;
      break;
    }

    if (seen.has(curr.slice(1).toString())) {
      continue;
    }
    seen.add(curr.slice(1).toString());

    if (n < 3 && !(dr == 0 && dc == 0)) {
      let [nr, nc] = [r + dr, c + dc];
      if (nr >= 0 && nr < grid.length && nc >= 0 && nc < grid[0].length) {
        heappush(hq, [hl + grid[nr][nc], nr, nc, dr, dc, n + 1]);
      }
    }

    for (let nd of [
      [0, 1],
      [1, 0],
      [0, -1],
      [-1, 0],
    ]) {
      let [ndr, ndc] = nd;
      if ((ndr == dr && ndc == dc) || (ndr == -dr && ndc == -dc)) {
        continue;
      }
      let [nr, nc] = [r + ndr, c + ndc];
      if (nr >= 0 && nr < grid.length && nc >= 0 && nc < grid[0].length) {
        heappush(hq, [hl + grid[nr][nc], nr, nc, ndr, ndc, 1]);
      }
    }
  }

  return total;
}
