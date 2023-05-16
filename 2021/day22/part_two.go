package main

func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func countCubes(ops []Operation) int {
	res := 0
	var st int

	for _, op := range ops {
		if op.type_op == "on" {
			st = 1
		} else {
			st = -1
		}
		res += st * AbsInt(
			op.max_x-op.min_x+1,
		) * AbsInt(
			op.max_y-op.min_y+1,
		) * AbsInt(
			op.max_z-op.min_z+1,
		)
	}
	return res
}

func part2(input string) int {
	operations := prepareData(input)
	tot_ops := make([]Operation, 0)

	for _, op1 := range operations {
		tmp_add := make([]Operation, 0)
		for _, op2 := range tot_ops {
			// for each compute intersection with previous ones
			if op2.max_x < op1.min_x || op1.max_x < op2.min_x || op2.max_y < op1.min_y ||
				op1.max_y < op2.min_y ||
				op2.max_z < op1.min_z ||
				op1.max_z < op2.min_z {
				continue
			}

			cx_min := MaxInt(op1.min_x, op2.min_x)
			cx_max := MinInt(op1.max_x, op2.max_x)
			cy_min := MaxInt(op1.min_y, op2.min_y)
			cy_max := MinInt(op1.max_y, op2.max_y)
			cz_min := MaxInt(op1.min_z, op2.min_z)
			cz_max := MinInt(op1.max_z, op2.max_z)
			var new_type string
			if op2.type_op == "on" {
				new_type = "off"
			} else {
				new_type = "on"
			}
			tmp_add = append(
				tmp_add,
				Operation{new_type, cx_min, cx_max, cy_min, cy_max, cz_min, cz_max},
			)
		}
		if op1.type_op == "on" {
			tot_ops = append(tot_ops, op1)
		}
		tot_ops = append(tot_ops, tmp_add...)
	}
	return countCubes(tot_ops)
}
