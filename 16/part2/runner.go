package aoc

import (
	"log"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	strm := stream{
		encoded: data[0],
	}
	return parsePacket(&strm)
}

type stream struct {
	encoded string
	pointer int
	buffer  string
}

func (f *stream) read(num int) string {
	for len(f.buffer) < num {
		if f.pointer >= len(f.encoded) {
			return ""
		}
		f.buffer += convert(f.encoded[f.pointer])
		f.pointer++
	}
	out := string(f.buffer[:num])
	f.buffer = f.buffer[num:]
	// fmt.Println(out)
	return out
}

func (f *stream) position() int {
	return f.pointer*4 - len(f.buffer)
}

func convert(char byte) string {
	var out string
	switch char {
	case '0':
		out = "0000"
	case '1':
		out = "0001"
	case '2':
		out = "0010"
	case '3':
		out = "0011"
	case '4':
		out = "0100"
	case '5':
		out = "0101"
	case '6':
		out = "0110"
	case '7':
		out = "0111"
	case '8':
		out = "1000"
	case '9':
		out = "1001"
	case 'A':
		out = "1010"
	case 'B':
		out = "1011"
	case 'C':
		out = "1100"
	case 'D':
		out = "1101"
	case 'E':
		out = "1110"
	case 'F':
		out = "1111"
	}
	return out
}

func parsePacket(strm *stream) int {
	strm.read(3)
	// Version parse
	s := strm.read(3)
	id, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal("id parse ", err)
	}
	if id == 4 {
		var newNum string
		for {
			s = strm.read(5)
			newNum += s[1:]
			if s[0] == '0' {
				break
			}
		}
		num, err := strconv.ParseInt(newNum, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		return int(num)
	}
	var numbers []int
	s = strm.read(1)
	if s == "0" {
		// 15-bit number for bits
		s = strm.read(15)
		length, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		start := strm.position()
		for strm.position()-start < int(length)-1 {
			numbers = append(numbers, parsePacket(strm))
		}
	} else {
		// 11-bit number for subpackets
		s = strm.read(11)
		pkts, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			log.Fatal("11 bit ", err)
		}
		for i := 0; i < int(pkts); i++ {
			numbers = append(numbers, parsePacket(strm))
		}
	}
	var count int
	switch id {
	case 0:
		for _, v := range numbers {
			count += v
		}
	case 1:
		count = 1
		for _, v := range numbers {
			count *= v
		}
	case 2:
		count = numbers[0]
		for _, v := range numbers {
			if v < count {
				count = v
			}
		}
	case 3:
		count = numbers[0]
		for _, v := range numbers {
			if v > count {
				count = v
			}
		}
	case 5:
		if numbers[0] > numbers[1] {
			count = 1
		}
	case 6:
		if numbers[0] < numbers[1] {
			count = 1
		}
	case 7:
		if numbers[0] == numbers[1] {
			count = 1
		}
	}
	return count
}
