package main

func part2(input string) int {
	bits := prepareData(input)
	val := HexToBit(bits)

	_, res, _ := handlePacket(val)
	return res
}
