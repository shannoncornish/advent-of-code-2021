package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type bitReader struct {
	s string

	position int
}

func newReader(s string) *bitReader {
	return &bitReader{
		s:        s,
		position: 0,
	}
}

func (br *bitReader) ReadBits(bits int) (int, int) {
	n := br.position + bits
	s := br.s[br.position:n]
	br.position = n

	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(i), bits
}

func (br *bitReader) ReadLiteral() (int, int) {
	var builder strings.Builder
	for {
		n := br.position + 5
		group := br.s[br.position:n]
		br.position = n

		builder.WriteString(group[1:])

		if last := group[0] == '0'; last {
			break
		}
	}

	i, err := strconv.ParseInt(builder.String(), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(i), (builder.Len() / 4) * 5
}

type packet struct {
	version int
	typeID  int

	literal int

	lengthTypeID int

	subpackets []packet
}

func readPacket(br *bitReader) (p packet, n int) {
	v, vn := br.ReadBits(3)
	n += vn

	p.version = v

	v, vn = br.ReadBits(3)
	n += vn

	p.typeID = v

	if p.typeID == 4 {
		v, vn = br.ReadLiteral()
		n += vn

		p.literal = v

	} else {
		v, vn := br.ReadBits(1)
		n += vn

		p.lengthTypeID = v
		if p.lengthTypeID == 0 {
			subpacketsBits, vn := br.ReadBits(15)
			n += vn

			for subpacketsBits > 0 {
				subpacket, vn := readPacket(br)
				n += vn
				subpacketsBits -= vn

				p.subpackets = append(p.subpackets, subpacket)
			}

		} else {
			subpacketsLen, vn := br.ReadBits(11)
			n += vn

			for i := 0; i < subpacketsLen; i++ {
				subpacket, vn := readPacket(br)
				n += vn

				p.subpackets = append(p.subpackets, subpacket)
			}
		}
	}

	return p, n
}

func sumVersions(p packet) int {
	sum := p.version
	for _, subpacket := range p.subpackets {
		sum += sumVersions(subpacket)
	}

	return sum
}

const (
	Sum = iota
	Product
	Minimum
	Maximum
	Literal
	GreaterThan
	LessThan
	EqualTo
)

func min(values []int) int {
	min := values[0]
	for _, value := range values[1:] {
		if value < min {
			min = value
		}
	}

	return min
}

func max(values []int) int {
	max := values[0]
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

func evaluate(p packet) int {
	switch p.typeID {
	case Sum:
		sum := 0
		for _, subpacket := range p.subpackets {
			sum += evaluate(subpacket)
		}
		return sum
	case Product:
		product := 1
		for _, subpacket := range p.subpackets {
			product *= evaluate(subpacket)
		}
		return product
	case Minimum:
		values := make([]int, len(p.subpackets))
		for i, subpacket := range p.subpackets {
			values[i] = evaluate(subpacket)
		}
		return min(values)
	case Maximum:
		values := make([]int, len(p.subpackets))
		for i, subpacket := range p.subpackets {
			values[i] = evaluate(subpacket)
		}
		return max(values)
	case Literal:
		return p.literal
	case GreaterThan:
		first, second := evaluate(p.subpackets[0]), evaluate(p.subpackets[1])
		if first > second {
			return 1
		}
		return 0
	case LessThan:
		first, second := evaluate(p.subpackets[0]), evaluate(p.subpackets[1])
		if first < second {
			return 1
		}
		return 0
	case EqualTo:
		first, second := evaluate(p.subpackets[0]), evaluate(p.subpackets[1])
		if first == second {
			return 1
		}
		return 0
	default:
		log.Fatalf("Unsupported packet type: %d\n", p.typeID)
		return -1
	}
}

func main() {
	r := newReader(parseInput())
	p, _ := readPacket(r)

	result := evaluate(p)

	fmt.Printf("The result of evaluating the expression: %d\n", result)

	// Answer: 2223947372407
}

func parseInput() string {
	name := "input.txt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	log.Printf("Input file: %s\n", name)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var builder strings.Builder
	for _, b := range bytes {
		var offset byte
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			offset = 48
			break
		case 'A', 'B', 'C', 'D', 'E', 'F':
			offset = 55
			break
		}

		builder.WriteString(fmt.Sprintf("%04b", b-offset))
	}

	return builder.String()
}
