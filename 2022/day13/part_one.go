package main

import (
	"os"
	"strconv"
	"strings"
)

type tree struct {
	valueLeaf int
	elements  []*tree
	father    *tree
}

func readTree(input string) tree {
	root := tree{-1, []*tree{}, nil}
	temp := &root

	var number string

	for _, chr := range input {
    switch chr {
    case '[' :
      tmpTree := tree{-1, []*tree{}, temp}
      temp.elements = append(temp.elements, &tmpTree)
      temp = &tmpTree
    case ']':
      if len(number) > 0 {
        num, _ := strconv.Atoi(number)
        temp.valueLeaf = num
        number = ""
      }
      temp = temp.father
    case ',':
      if len(number) > 0 {
        num, _ := strconv.Atoi(number)
        temp.valueLeaf = num
        number = ""
      }
      temp = temp.father
      tmpTree := tree{-1, []*tree{}, temp}
      temp.elements = append(temp.elements, &tmpTree)
      temp = &tmpTree
    default:
      number += string(chr)
    }	
	}
	return root
}

func checkOrder(t1 tree, t2 tree) int {
  switch {
  case len(t1.elements) == 0 && len(t2.elements) == 0:
		if t1.valueLeaf > t2.valueLeaf {
			return -1
		} else if t1.valueLeaf == t2.valueLeaf {
			return 0
		} 
    return 1
	
  case t1.valueLeaf >= 0:
    return checkOrder(tree{-1, []*tree{&t1}, nil}, t2)
  case t2.valueLeaf >= 0:
		return checkOrder(t1, tree{-1, []*tree{&t2}, nil})
  default:
    var i int
		for i = 0; i < len(t1.elements) && i < len(t2.elements); i++ {
			ordered := checkOrder(*t1.elements[i], *t2.elements[i])
			if ordered != 0 {
				return ordered
			}
		}
		if i < len(t1.elements) {
			return -1
		} else if i < len(t2.elements) {
			return 1
		}
	}
	return 0
}

func part1() int {
	dat, _ := os.ReadFile("./day13-input")
	s := strings.Split(string(dat), "\n\n")
  var order_num int
  for i, pair := range s {
    pair_split := strings.Split(pair, "\n")
    p1, p2 := readTree(pair_split[0]), readTree(pair_split[1])
    if checkOrder(p1, p2) == 1 {
      order_num += i + 1
    }
  }
  return order_num
}
