package sixteen

import (
	"2021/Utilities"
	"fmt"
	"math"
	"strconv"
)

func sum(bits []int) int {
	total := 0
	for _, b := range bits {
		total += b
	}
	return total
}

func product(bits []int) int {
	total := 1
	for _, b := range bits {
		total *= b
	}
	return total
}

func min(total []int) int {
	m := math.MaxInt
	for _, n := range total {
		if n < m {
			m = n
		}
	}
	return m
}

func max(total []int) int {
	m := math.MinInt
	for _, n := range total {
		if n > m {
			m = n
		}
	}
	return m
}

func calculate(total []int, t int) int {
	switch t {
	case 0:
		return sum(total)
	case 1:
		return product(total)
	case 2:
		return min(total)
	case 3:
		return max(total)
	case 5:
		if total[0] > total[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if total[0] < total[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if total[0] == total[1] {
			return 1
		} else {
			return 0
		}
	default:
		panic("Type not correct " + strconv.Itoa(t))
	}
}

func decodeBits2(bits []int) (int, int) {
	if len(bits) == 0 {
		return 0, 0
	}
	i := 3
	t := bitsToDec(bits[i : i+3])
	i += 3
	switch t {
	case 4:
		number := make([]int, 0)
		number = append(number, bits[i+1:i+5]...)
		for bits[i] == 1 {
			i += 5
			number = append(number, bits[i+1:i+5]...)
		}
		return bitsToDec(number), i + 5
	default:
		length := bits[i]
		i += 1
		total := make([]int, 0)
		if length == 0 {
			lenSubPackets := bitsToDec(bits[i : i+15])
			i += 15
			totalLength := i + lenSubPackets
			for i < totalLength {
				value, position := decodeBits2(bits[i:])
				i += position
				total = append(total, value)
			}
		} else {
			numberSubPackets := bitsToDec(bits[i : i+11])
			i += 11
			for x := 0; x < numberSubPackets; x++ {
				value, position := decodeBits2(bits[i:])
				total = append(total, value)
				i += position
			}
		}
		return calculate(total, t), i
	}
}

func Second(splittedStrings []string) int {
	bits := hexToBits(splittedStrings[0])
	value, _ := decodeBits2(bits)
	return value
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
