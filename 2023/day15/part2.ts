// Part 2 for day 15 of 2023

export function part2(input: string): number {
  let steps = input
    .split(",")
    .filter((c) => c != "")
    .map((val) => val.replace(/(\r\n|\n|\r)/gm, ""));

  let boxes: { [index: number]: string[][] } = {};

  for (let i = 0; i < steps.length; i++) {
    let step = steps[i];

    let label = step.includes("=") ? step.split("=")[0] : step.split("-")[0];
    let boxLabel = label
      .split("")
      .reduce((x, y) => ((x + y.charCodeAt(0)) * 17) % 256, 0);
    let boxType = step.includes("=") ? "add" : "del";

    if (boxType == "add") {
      if (!boxes[boxLabel]) {
        // if does not exist yet
        boxes[boxLabel] = [step.split("=")];
      } else {
        // if box exists but not the label, just push
        let exists = false;
        for (let i = 0; i < boxes[boxLabel].length; i++) {
          if (boxes[boxLabel][i][0] == label) {
            boxes[boxLabel][i][1] = step.split("=")[1];
            exists = true;
            break;
          }
        }
        if (!exists) {
          // if the label does not exist in the box we add it
          boxes[boxLabel].push(step.split("="));
        }
      }
    } else if (boxType == "del" && boxes[boxLabel]) {
      for (let i = 0; i < boxes[boxLabel].length; i++) {
        if (boxes[boxLabel][i][0] == label) {
          boxes[boxLabel].splice(i, 1);
          break;
        }
      }
    }
  }

  let total = 0;
  for (let box in boxes) {
    for (let i = 0; i < boxes[box].length; i++) {
      let tmp = (parseInt(box) + 1) * (i + 1) * parseInt(boxes[box][i][1]);
      total += tmp;
    }
  }

  return total;
}
