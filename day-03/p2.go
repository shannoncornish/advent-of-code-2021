package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://stackoverflow.com/questions/23192262/how-would-you-set-and-clear-a-single-bit-in-go

func main() {
	values, size := parseInput()

	set, unset := split(values, size-1)

	var mostCommon, leastCommon []uint
	if len(set) > len(unset) {
		mostCommon = set
		leastCommon = unset
	} else {
		mostCommon = unset
		leastCommon = set
	}

	oxygen := reduceMostCommon(mostCommon, size-2)
	co2 := reduceLeastCommon(leastCommon, size-2)

	log.Printf("Oxygen: %08b\n", oxygen)
	log.Printf("CO2: %08b\n", co2)

	fmt.Printf("Oxygen: %d, CO2: %d, Life Support: %d\n", oxygen, co2, oxygen*co2)

	// Answer: 3414905
}

func reduceMostCommon(values []uint, pos int) uint {
	if len(values) == 1 {
		return values[0]
	}

	set, unset := split(values, pos)
	if len(set) >= len(unset) {
		return reduceMostCommon(set, pos-1)
	} else {
		return reduceMostCommon(unset, pos-1)
	}
}

func reduceLeastCommon(values []uint, pos int) uint {
	if len(values) == 1 {
		return values[0]
	}

	set, unset := split(values, pos)
	if len(set) >= len(unset) && len(unset) > 0 {
		return reduceLeastCommon(unset, pos-1)
	} else {
		return reduceLeastCommon(set, pos-1)
	}
}

func split(values []uint, pos int) (set, unset []uint) {
	for i := range values {
		value := values[i] & (1 << pos)
		if value > 0 {
			set = append(set, values[i])
		} else {
			unset = append(unset, values[i])
		}
	}

	return set, unset
}

func parseInput() ([]uint, int) {
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

	scanner := bufio.NewScanner(f)

	var size int
	var values []uint
	for scanner.Scan() {
		text := scanner.Text()
		if size == 0 {
			size = len(text)
		}

		value, err := strconv.ParseUint(text, 2, 32)
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, uint(value))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values, size
}
