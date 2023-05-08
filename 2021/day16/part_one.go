package main

import (
	"fmt"
	"strconv"
)

func HexToBit(input string) string {
	dic := map[string]string{
		"0": "0000", "1": "0001", "2": "0010", "3": "0011", "4": "0100",
		"5": "0101", "6": "0110", "7": "0111", "8": "1000", "9": "1001", "A": "1010",
		"B": "1011", "C": "1100", "D": "1101", "E": "1110", "F": "1111",
	}
	tmp := ""

	for _, num := range input {
		tmp += dic[string(num)]
	}
	return tmp
}

func prepareData(input string) string {
	return input[:len(input)-1]
}

func handlePacket(pack string) (versionTotal int64, expressionValue int, bitsRead int) {
	version, _ := strconv.ParseInt(pack[0:3], 2, 64)
	versionTotal += version

	typeID, _ := strconv.ParseInt(pack[3:6], 2, 64)

	read := 6 // version and typeID

	switch typeID {
	case 4: // literal value
		// parse 5 bits at a time
		var bits string
		for i := 6; i < len(pack); i += 5 {
			fiveBits := pack[i : i+5]
			read += 5

			bits += fiveBits[1:]

			if fiveBits[0] == '0' {
				break
			}
		}

		decimalVal, _ := strconv.ParseInt(bits, 2, 64)
		return version, int(decimalVal), read
	default: // operator types?
		// contains one or more packets
		lengthTypeID := pack[6:7]
		read++
		var bitsToRead int
		switch lengthTypeID {
		case "0":
			// next 15 bits == total length in bits for REST of subpackets
			bitsToRead = 15
		case "1":
			// next 11 bits == NUMBER of subpackets
			bitsToRead = 11
		}

		rawLength := pack[7 : 7+bitsToRead]
		read += bitsToRead
		length, _ := strconv.ParseInt(rawLength, 2, 64)

		// followed by the subpackets themselves
		var subPacketExpressionValues []int
		switch lengthTypeID {
		case "0":
			// next 15 bits == total length in bits for REST of subpackets
			for length > 0 {
				// continue reading until length of bits are read
				ver, expVal, n := handlePacket(pack[read:])
				read += n
				length -= int64(n)
				subPacketExpressionValues = append(subPacketExpressionValues, expVal)
				versionTotal += ver
			}
		case "1":
			// next 11 bits == NUMBER of subpackets
			for length > 0 {
				// continue reading until number of packets are read
				ver, expVal, n := handlePacket(pack[read:])
				read += n
				length-- // reduce length by 1 (ie one packet read)
				subPacketExpressionValues = append(subPacketExpressionValues, expVal)
				versionTotal += ver
			}
		}

		switch typeID {
		// note: case 0 already handled above, literal value
		case 0: // sum
			return versionTotal, SumIntSlice(subPacketExpressionValues), read
		case 1: // product
			return versionTotal, MultiplyIntSlice(subPacketExpressionValues), read
		case 2: // min
			return versionTotal, MinInt(subPacketExpressionValues...), read
		case 3: // max
			return versionTotal, MaxInt(subPacketExpressionValues...), read
			// 4 is literal...
		case 5: // greater than (first subpacket > second, will always have exactly 2)
			var ans int
			if subPacketExpressionValues[0] > subPacketExpressionValues[1] {
				ans = 1 // otherwise int zero val works
			}
			return versionTotal, ans, read
		case 6: // less than (opposite)
			var ans int
			if subPacketExpressionValues[0] < subPacketExpressionValues[1] {
				ans = 1 // otherwise int zero val works
			}
			return versionTotal, ans, read
		case 7: // equal to
			var ans int
			if subPacketExpressionValues[0] == subPacketExpressionValues[1] {
				ans = 1 // otherwise int zero val works
			}
			return versionTotal, ans, read
		default:
			panic(fmt.Sprintf("unknown typeID: %d", typeID))
		}
	}
}

func SumIntSlice(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func MultiplyIntSlice(nums []int) int {
	product := 1
	for _, n := range nums {
		product *= n
	}
	return product
}

func MaxInt(nums ...int) int {
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

func MinInt(nums ...int) int {
	minNum := nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}

func part1(input string) int64 {
	bits := prepareData(input)
	val := HexToBit(bits)

	version, _, _ := handlePacket(val)
	return version
}
