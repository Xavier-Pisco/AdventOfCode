package sixteen

import (
	"2021/Utilities"
	"fmt"
	"math"
)

func Second(splittedStrings []string) int {
	return 0
}

func hexToBits(s string) []int {
	bits := make([]int, 0)
	for _, c := range s {
		switch c {
		case '0':
			bits = append(bits, 0, 0, 0, 0)
		case '1':
			bits = append(bits, 0, 0, 0, 1)
		case '2':
			bits = append(bits, 0, 0, 1, 0)
		case '3':
			bits = append(bits, 0, 0, 1, 1)
		case '4':
			bits = append(bits, 0, 1, 0, 0)
		case '5':
			bits = append(bits, 0, 1, 0, 1)
		case '6':
			bits = append(bits, 0, 1, 1, 0)
		case '7':
			bits = append(bits, 0, 1, 1, 1)
		case '8':
			bits = append(bits, 1, 0, 0, 0)
		case '9':
			bits = append(bits, 1, 0, 0, 1)
		case 'A':
			bits = append(bits, 1, 0, 1, 0)
		case 'B':
			bits = append(bits, 1, 0, 1, 1)
		case 'C':
			bits = append(bits, 1, 1, 0, 0)
		case 'D':
			bits = append(bits, 1, 1, 0, 1)
		case 'E':
			bits = append(bits, 1, 1, 1, 0)
		case 'F':
			bits = append(bits, 1, 1, 1, 1)
		}
	}
	return bits
}

func bitsToDec(bits []int) int {
	number := 0
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			number += int(math.Pow(2, float64(len(bits)-1-i)))
		}
	}
	return number
}

func decodeBits(bits []int) (int, int) {
	if len(bits) == 0 {
		return 0, 0
	}
	i := 0
	version := bitsToDec(bits[i : i+3])
	i += 3
	t := bitsToDec(bits[i : i+3])
	i += 3
	switch t {
	case 4:
		number := bitsToDec(bits[i : i+5])
		i += 5
		for number > 15 {
			number = bitsToDec(bits[i : i+5])
			i += 5
		}
		return version, i
	default:
		length := bits[i]
		i += 1
		total := 0
		if length == 0 {
			lenSubPackets := bitsToDec(bits[i : i+15])
			i += 15
			totalLength := i + lenSubPackets
			for i < totalLength {
				value, position := decodeBits(bits[i:])
				i += position
				total += value
			}
		} else {
			numberSubPackets := bitsToDec(bits[i : i+11])
			i += 11
			for x := 0; x < numberSubPackets; x++ {
				value, position := decodeBits(bits[i:])
				total += value
				i += position
			}
		}
		return version + total, i
	}
}

func First(splittedStrings []string) int {
	bits := hexToBits(splittedStrings[0])
	value, _ := decodeBits(bits)
	return value
}

func Solve() {
	splittedStrings := Utilities.Read("16/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
